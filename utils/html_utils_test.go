package utils

import (
	"testing"

	"golang.org/x/net/html"
)

var utils UtilProvider = &Utils{}

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
			htmlVersion := utils.ExtractHTMLVersion(test.htmlStr)

			if htmlVersion != test.expected {
				t.Errorf("Expected HTML version '%s', got '%s'", test.expected, htmlVersion)
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
			node, err := utils.ParseHTML(test.html)

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
			res := utils.ExtractAttribute(test.node, test.attr)

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
			res := utils.HasLoginForm(test.node)

			if res != test.expected {
				t.Errorf("Expected '%t', got '%t'", test.expected, res)
			}
		})
	}
}

func TestExtractTitle(t *testing.T) {
	tests := []struct {
		name     string
		node     *html.Node
		expected string
	}{
		{
			name: "Has title",
			node: &html.Node{
				FirstChild: &html.Node{
					Type: html.TextNode,
					Data: "Test title",
				},
			},
			expected: "Test title",
		},
		{
			name:     "Title does not exist",
			node:     &html.Node{},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := utils.ExtractTitle(test.node)

			if res != test.expected {
				t.Errorf("Expected title '%s', got '%s'", test.expected, res)
			}
		})
	}
}
