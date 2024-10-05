package utils

import (
	"html/template"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// RenderTemplate renders the specified template with the given data
func (u *Utils) RenderTemplate(w http.ResponseWriter, r *http.Request, templatePath string, data any) error {
	t := template.Must(template.ParseFiles(templatePath))

	if err := t.Execute(w, data); err != nil {
		return err
	}
	return nil
}

// ParseHTML parses the HTML content and return as a HTML node tree
func (u *Utils) ParseHTML(pageHTML string) (*html.Node, error) {
	doc, err := html.Parse(strings.NewReader(pageHTML))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// ExtractAttribute returns the value of the given attribute from the specified HTML node
func (u *Utils) ExtractAttribute(n *html.Node, attr string) string {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val
		}
	}
	return ""
}

// HasLoginForm checks the given HTML node has a login form
func (u *Utils) HasLoginForm(n *html.Node) bool {
	attrVal := u.ExtractAttribute(n, "type")

	return strings.ToLower(attrVal) == "password"
}

// ExtractTitle returns the value of the title element in the specified HTML node
func (u *Utils) ExtractTitle(n *html.Node) string {
	if n.FirstChild != nil {
		return n.FirstChild.Data
	}
	return ""
}

// ExtractHTMLVersion returns the version of the HTML of the given HTML content
func (u *Utils) ExtractHTMLVersion(htmlContent string) string {
	content := strings.ToLower(htmlContent)
	content = strings.Trim(content, "\n")

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
