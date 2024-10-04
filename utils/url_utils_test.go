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
			expectedError: errors.New(fmt.Sprintf("unexpected status code received. status code: %d", http.StatusNotFound)),
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

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "Network error" {
				// Simulate a network error by not calling the server
				body, err := FetchURL("http://non-existent-url")
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

			body, err := FetchURL(server.URL)

			if body != test.expectedBody {
				t.Errorf("Expected body '%s', got '%s'", test.expectedBody, body)
			}

			if (err != nil && test.expectedError == nil) || (err == nil && test.expectedError != nil) {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			} else if err != nil && test.expectedError != nil && err.Error() != test.expectedError.Error() {
				t.Errorf("expected error %q, got %q", test.expectedError.Error(), err.Error())
			}
		})
	}
}
