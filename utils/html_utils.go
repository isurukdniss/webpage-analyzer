package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, templatePath string, data any) error { //TODO: test the any type

	t := template.Must(template.ParseFiles(templatePath))

	if err := t.Execute(w, data); err != nil {
		return err
	}
	return nil
}

func ParseHTML(pageHtml string) (*html.Node, error) {
	doc, err := html.Parse(strings.NewReader(pageHtml))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func ExtractAttribute(n *html.Node, attr string) string {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val
		}
	}
	return ""
}

func HasLoginForm(n *html.Node) bool {
	attrVal := ExtractAttribute(n, "type")

	return strings.ToLower(attrVal) == "password"
}

func ExtractTitle(n *html.Node) string {
	var title string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			title = c.Data
		}
	}
	return title
}

func IsLinkAccessible(link string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Head(link)
	if err != nil {
		return false
	}
	defer res.Body.Close()

	return res.StatusCode < 400
}

func ExtractHTMLVersion(htmlContent string) string {
	content := strings.ToLower(htmlContent)

	if strings.HasPrefix(content, "<!doctype html>") {
		return "HTML 5"
	} else if strings.Contains(content, `"-//w3c//dtd html 4.01//en"`) {
		return "HTML 4.01"
	} else if strings.Contains(content, `"-//w3c//dtd xhtml 1.0 strict//en"`) {
		return "XHTML 1.0 Strict"
	} else if strings.Contains(content, `"-//w3c//dtd xhtml 1.0 transitional//en"`) {
		return "XHTML 1.0 Transitional"
	} else if strings.Contains(content, `"-//w3c//dtd xhtml 1.1//en"`) {
		return "XHTML 1.1"
	}

	return "Unknown"
}

func IsInternalLink(baseURL string, targetURL string) (bool, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return false, fmt.Errorf("invalid base URL: %w", err)
	}

	target, err := url.Parse(targetURL)
	if err != nil {
		return false, fmt.Errorf("invalid target URL: %w", err)
	}

	if !target.IsAbs() {
		return true, nil
	}

	hasSameHost := strings.EqualFold(base.Host, target.Host)

	return hasSameHost, nil
}
