package utils

import (
	"net/http"

	"golang.org/x/net/html"
)

type UtilProvider interface {
	RenderTemplate(w http.ResponseWriter, r *http.Request, templatePath string, data any) error
	ExtractTitle(n *html.Node) string
	HasLoginForm(n *html.Node) bool
	ExtractAttribute(n *html.Node, attr string) string
	IsLinkAccessible(link string) bool
	IsInternalLink(baseURL string, targetURL string) (bool, error)
	ExtractHTMLVersion(htmlContent string) string
	ParseHTML(pageHtml string) (*html.Node, error)
	FetchURL(url string) (string, error)
}

type Utils struct{}
