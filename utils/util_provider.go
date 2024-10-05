package utils

import (
	"net/http"

	"golang.org/x/net/html"
)

// UtilProvider defines a set of utility methods used for rendering templates, extracting
// data from HTML nodes, checking link accessibility, and fetching URLs.
type UtilProvider interface {
	RenderTemplate(w http.ResponseWriter, r *http.Request, templatePath string, data any) error
	ExtractTitle(n *html.Node) string
	HasLoginForm(n *html.Node) bool
	ExtractAttribute(n *html.Node, attr string) string
	IsLinkAccessible(link string) bool
	IsInternalLink(baseURL string, targetURL string) bool
	ExtractHTMLVersion(htmlContent string) string
	ParseHTML(pageHTML string) (*html.Node, error)
	FetchURL(url string) (string, error)
}

// Utils provides utility functions for handling common operations
// like rendering templates, parsing HTML, and URL-related tasks.
type Utils struct{}
