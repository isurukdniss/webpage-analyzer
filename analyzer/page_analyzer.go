package analyzer

import (
	"fmt"
	"strings"
	"sync"

	"github.com/isurukdniss/webpage-analyzer/utils"

	"golang.org/x/net/html"
)

var utilsInstance utils.UtilProvider = &utils.Utils{}

// Result represents the webpage analyzer output data structure
type Result struct {
	HTMLVersion        string
	Title              string
	HeadingsCount      map[string]int
	InternalLinksCount int
	ExternalLinksCount int
	InAccessibleLinks  int
	HasLoginForm       bool
	ErrorMessage       string
	ExternalLinks      []string
}

// PageAnalyzer defines the interface for analyzing a webpage based on its URL
type PageAnalyzer interface {
	Analyze(pageURL string) *Result
}

// Analyzer provides function to analyze the HTML content of a given URL
type Analyzer struct{}

// Analyze function analyzes the HTML content of the website of a given URL
func (a *Analyzer) Analyze(pageURL string) *Result {
	res := &Result{
		HeadingsCount: make(map[string]int),
	}
	// To track the visited links
	visited := make(map[string]bool)

	body, err := utilsInstance.FetchURL(pageURL)
	if err != nil {
		res.ErrorMessage = handleErrorMsg(err)
	}

	doc, err := utilsInstance.ParseHTML(body)
	if err != nil {
		res.ErrorMessage = handleErrorMsg(err)
	}

	res.HTMLVersion = utilsInstance.ExtractHTMLVersion(body)

	analyzeDoc(doc, pageURL, visited, res)

	externalLinks := res.ExternalLinks
	// Inaccessible links check is performed only for external links
	inAccessibleLinksCount := getInaccessibleLinksCount(externalLinks)
	res.InAccessibleLinks = inAccessibleLinksCount

	return res

}

func analyzeDoc(n *html.Node, baseURL string, visited map[string]bool, res *Result) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "title":
			// some webpage's html may contain multiple <title> tags. eg. inside svg tags.
			if res.Title == "" {
				res.Title = utilsInstance.ExtractTitle(n)
			}
		case "h1", "h2", "h3", "h4", "h5", "h6":
			res.HeadingsCount[n.Data]++
		case "input":
			res.HasLoginForm = utilsInstance.HasLoginForm(n)
		case "a":
			link := utilsInstance.ExtractAttribute(n, "href")
			if !visited[link] {
				visited[link] = true
				if isInternal := utilsInstance.IsInternalLink(baseURL, link); isInternal {
					res.InternalLinksCount++
				} else {
					res.ExternalLinksCount++
					res.ExternalLinks = append(res.ExternalLinks, link)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		analyzeDoc(c, baseURL, visited, res)
	}
}

func getInaccessibleLinksCount(urlList []string) int {
	var wg sync.WaitGroup
	//mutex is used to lock the 'count' when it is updated by multiple goroutines
	var mu sync.Mutex
	var count int

	for _, link := range urlList {
		wg.Add(1)

		go func(link string) {
			defer wg.Done()
			if !utilsInstance.IsLinkAccessible(link) {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}(link)
	}
	wg.Wait()
	return count
}

func handleErrorMsg(err error) string {
	if strings.Contains(err.Error(), "invalid URI for request") {
		return "The provided URL is not valid. Please check the format and try again."
	}
	if err.Error() == "invalid URL: missing scheme or host" {
		return "The URL is missing a scheme (like 'http' or 'https') or a host. Please provide a complete URL."
	}
	if err.Error() == "unable to fetch the URL" {
		return "We were unable to fetch the requested URL. Please check your internet connection or the URL."
	}
	if strings.Contains(err.Error(), "status code") {
		split := strings.Split(err.Error(), " ")
		sc := split[len(split)-1]
		return fmt.Sprintf("The server returned a status code of %s. Please ensure you have the necessary permissions.", sc)
	}
	if err.Error() == "error reading the response body" {
		return "An error occurred while reading the response. Please try again later."
	}
	return "An unexpected error occurred. Please try again."
}
