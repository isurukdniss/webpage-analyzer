package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/html"
)

func TestExtractHTMLVersion(t *testing.T) {
	tests := []struct {
		name     string
		htmlStr  string
		expected string
	}{
		{
			name:     "Version: HTML 5",
			htmlStr:  "<!DOCTYPE html><html><head><title>Test</title></head><body></body></html>",
			expected: "HTML 5",
		},
		{
			name: "Version: HTML 4.01",
			htmlStr: `"<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
   						"http://www.w3.org/TR/html4/strict.dtd">
						<html>
							<head>
								<title>Test</title>
							</head>
							<body>
							</body>
						</html>"`,
			expected: "HTML 4.01",
		},
		{
			name: "Version: XHTML 1.0 Strict",
			htmlStr: `"<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    						"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
						<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
							<head>
								<title>Sample HTML 4.01 Document</title>
							</head>
							<body>
							</body>
						</html>"`,
			expected: "XHTML 1.0 Strict",
		},
		{
			name: "Version: XHTML 1.0 Transitional",
			htmlStr: `"<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" 
							"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
						<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
							<head>
								<title>Sample XHTML 1.0 Transitional Document</title>
							</head>
							<body>
							</body>
						</html>"`,
			expected: "XHTML 1.0 Transitional",
		},
		{
			name: "Version: XHTML 1.1",
			htmlStr: `"<?xml version="1.0" encoding="UTF-8"?>
						<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"
							"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
						<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
						<head>
								<title>Sample XHTML 1.0 Transitional Document</title>
							</head>
							<body>
							</body>
						</html>"`,
			expected: "XHTML 1.1",
		},
		{
			name:     "Unknown HTML version",
			htmlStr:  "",
			expected: "Unknown",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			htmlVersion := ExtractHTMLVersion(test.htmlStr)

			if htmlVersion != test.expected {
				t.Errorf("Expected HTML version '%s', got '%s'", test.expected, htmlVersion)
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
			isInternalLink, _ := IsInternalLink(test.baseURL, test.targetURL)

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

			isAccessible := IsLinkAccessible(server.URL)

			if isAccessible != test.expected {
				t.Errorf("Expected '%t', got '%t'", test.expected, isAccessible)
			}
		})
	}
}

func TestParseHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		hasError bool
	}{
		{
			name:     "Valid HTML",
			html:     "<html><head><title>Test</title></head><body><p>Test</p></body></html>",
			hasError: false,
		},
		{
			name:     "Invalid HTML",
			html:     "<html><head><title>Test</title><body><p>Test",
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			node, err := ParseHTML(test.html)

			if err != nil && !test.hasError {
				t.Error("Expected HTML parse no error but error returned")
			}

			if node == nil && !test.hasError {
				t.Error("Expected HTML parse an error but no error returned")
			}
		})
	}
}

func TestExtractAttribute(t *testing.T) {
	tests := []struct {
		name     string
		node     *html.Node
		attr     string
		expected string
	}{
		{
			name: "Contains the Attribute",
			node: &html.Node{
				Attr: []html.Attribute{
					{Key: "class", Val: "my-class"},
					{Key: "href", Val: "/test"},
				},
			},
			attr:     "href",
			expected: "/test",
		},
		{
			name: "Doesn't contain the Attribute",
			node: &html.Node{
				Attr: []html.Attribute{
					{Key: "class", Val: "my-class"},
					{Key: "href", Val: "/test"},
				},
			},
			attr:     "id",
			expected: "",
		},
		{
			name: "Empty Attribute",
			node: &html.Node{
				Attr: []html.Attribute{},
			},
			attr:     "id",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ExtractAttribute(test.node, test.attr)

			if res != test.expected {
				t.Errorf("Expected attribute value '%s', got '%s'", test.expected, res)
			}
		})
	}
}

func TestHasLoginForm(t *testing.T) {
	tests := []struct {
		name     string
		node     *html.Node
		expected bool
	}{
		{
			name: "Contains the password input",
			node: &html.Node{
				Attr: []html.Attribute{
					{Key: "class", Val: "my-class"},
					{Key: "type", Val: "password"},
				},
			},
			expected: true,
		},
		{
			name: "No password input",
			node: &html.Node{
				Attr: []html.Attribute{
					{Key: "class", Val: "my-class"},
					{Key: "id", Val: "my-id"},
				},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := HasLoginForm(test.node)

			if res != test.expected {
				t.Errorf("Expected '%t', got '%t'", test.expected, res)
			}
		})
	}
}
