package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchURL(t *testing.T) {

	tests := []struct {
		name          string
		statusCode    int
		responseBody  string
		expectedBody  string
		expectedError error
	}{
		{
			name:          "URL fetch success",
			statusCode:    http.StatusOK,
			responseBody:  "Success!",
			expectedBody:  "Success!",
			expectedError: nil,
		},
		{
			name:          "Unexpected Status Code",
			statusCode:    http.StatusNotFound,
			responseBody:  "",
			expectedBody:  "",
			expectedError: fmt.Errorf("unexpected status code: %d", http.StatusNotFound),
		},
		{
			name:          "Network error",
			statusCode:    0,
			responseBody:  "",
			expectedBody:  "",
			expectedError: errors.New("unable to fetch the URL"),
		},
		{
			name:          "Read response body error",
			statusCode:    http.StatusOK,
			responseBody:  "This is a test page.",
			expectedBody:  "This is a test page.",
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "Network error" {
				// Simulate a network error by not calling the server
				body, err := utils.FetchURL("http://non-existent-url")
				if body != test.expectedBody {
					t.Errorf("expected body %q, got %q", test.expectedBody, body)
				}
				if err == nil || err.Error() != test.expectedError.Error() {
					t.Errorf("expected error %q, got %v", test.expectedError.Error(), err)
				}
				return
			}

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.statusCode)
				io.WriteString(w, test.responseBody)
			}))
			defer server.Close()

			body, err := utils.FetchURL(server.URL)

			if body != test.expectedBody {
				t.Errorf("Expected body '%s', got '%s'", test.expectedBody, body)
			}

			if (err != nil && test.expectedError == nil) || (err == nil && test.expectedError != nil) {
				t.Errorf("Expected error %v, got %v", test.expectedError, err)
			} else if err != nil && test.expectedError != nil && err.Error() != test.expectedError.Error() {
				t.Errorf("Expected error %q, got %q", test.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestIsInternalLink(t *testing.T) {
	tests := []struct {
		name      string
		baseURL   string
		targetURL string
		expected  bool
	}{
		{
			name:      "Invalid base URL",
			baseURL:   ";;:::12abc",
			targetURL: "https://www.google.com/",
			expected:  false,
		},
		{
			name:      "Invalid target URL",
			baseURL:   "https://www.google.com/",
			targetURL: ";;:::12abc",
			expected:  false,
		},
		{
			name:      "Internal link",
			baseURL:   "https://www.google.com/",
			targetURL: "https://www.google.com/test",
			expected:  true,
		},
		{
			name:      "External link",
			baseURL:   "https://www.google.com/",
			targetURL: "https://www.yahoo.com/",
			expected:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isInternalLink := utils.IsInternalLink(test.baseURL, test.targetURL)

			if test.expected != isInternalLink {
				t.Errorf("Expected '%t', got '%t'", test.expected, isInternalLink)
			}
		})
	}
}

func TestIsLinkAccessible(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		expected   bool
	}{
		{
			name:       "Link Accessible",
			statusCode: http.StatusOK,
			expected:   true,
		},
		{
			name:       "Redirect Link",
			statusCode: http.StatusMovedPermanently,
			expected:   true,
		},
		{
			name:       "Inaccessible: 404 Not found",
			statusCode: http.StatusNotFound,
			expected:   false,
		},
		{
			name:       "Inaccessible: Internal Server Error",
			statusCode: http.StatusInternalServerError,
			expected:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.statusCode)
			}))
			defer server.Close()

			isAccessible := utils.IsLinkAccessible(server.URL)

			if isAccessible != test.expected {
				t.Errorf("Expected '%t', got '%t'", test.expected, isAccessible)
			}
		})
	}
}
