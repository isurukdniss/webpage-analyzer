package analyzer

import (
	"sync"

	"github.com/isurukdniss/webpage-analyzer/utils"

	"golang.org/x/net/html"
)

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
	body, err := utils.FetchURL(pageUrl)
	if err != nil {
		res.ErrorMessage = "Error analyzing the URL."

	}

	// parse html
	doc, err := utils.ParseHTML(body)
	if err != nil {
		res.ErrorMessage = "Error analyzing the URL."
	}

	res.HTMLVersion = utils.ExtractHTMLVersion(body)

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
			res.Title = utils.ExtractTitle(n)
		case "h1", "h2", "h3", "h4", "h5", "h6":
			res.HeadingsCount[n.Data]++
		case "input":

			res.HasLoginForm = utils.HasLoginForm(n)
		case "a":
			link := utils.ExtractAttribute(n, "href")
			if !visited[link] {
				visited[link] = true
				if isInternal, _ := utils.IsInternalLink(baseURL, link); isInternal { // Handle error
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
			if !utils.IsLinkAccessible(link) {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}(link)
	}
	wg.Wait()
	return count
}
