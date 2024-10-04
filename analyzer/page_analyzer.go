package analyzer

import (
	"fmt"
	"sync"

	"github.com/isurukdniss/webpage-analyzer/utils"

	"golang.org/x/net/html"
)

var utilsInstance utils.UtilProvider = &utils.Utils{}

type AnalyzerResult struct {
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

func Analyze(pageUrl string) *AnalyzerResult {
	res := &AnalyzerResult{
		HeadingsCount: make(map[string]int),
	}
	visited := make(map[string]bool)

	// fetch url
	body, err := utilsInstance.FetchURL(pageUrl)
	if err != nil {
		res.ErrorMessage = fmt.Sprintf("Error analyzing the URL: %v", err)

	}

	// parse html
	doc, err := utilsInstance.ParseHTML(body)
	if err != nil {
		res.ErrorMessage = fmt.Sprintf("Error analyzing the URL: %v", err)
	}

	res.HTMLVersion = utilsInstance.ExtractHTMLVersion(body)

	analyzeDoc(doc, pageUrl, visited, res)

	externalLinks := res.ExternalLinks
	inAccessibleLinksCount := getInaccessibleLinksCount(externalLinks)
	res.InAccessibleLinks = inAccessibleLinksCount

	return res

}

func analyzeDoc(n *html.Node, baseURL string, visited map[string]bool, res *AnalyzerResult) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "title":
			res.Title = utilsInstance.ExtractTitle(n)
		case "h1", "h2", "h3", "h4", "h5", "h6":
			res.HeadingsCount[n.Data]++
		case "input":

			res.HasLoginForm = utilsInstance.HasLoginForm(n)
		case "a":
			link := utilsInstance.ExtractAttribute(n, "href")
			if !visited[link] {
				visited[link] = true
				if isInternal, _ := utilsInstance.IsInternalLink(baseURL, link); isInternal { // Handle error
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
