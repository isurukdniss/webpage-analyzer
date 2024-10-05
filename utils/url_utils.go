package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FetchURL fetches the specified URL and return the HTML content as a string
func (u *Utils) FetchURL(rawURL string) (string, error) {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", errors.New("invalid URL: missing scheme or host")
	}

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return "", errors.New("unable to fetch the URL")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("unexpected status code: %d", resp.StatusCode)
		return "", errors.New(errMsg)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("error reading the response body")
	}

	return string(body), nil
}

// IsInternalLink checks whether the given targetURL is internal to the baseURL
func (u *Utils) IsInternalLink(baseURL string, targetURL string) bool {
	base, err := url.Parse(baseURL)
	if err != nil {
		log.Println(err)
		return false
	}

	target, err := url.Parse(targetURL)
	if err != nil {
		log.Println(err)
		return false
	}

	if !target.IsAbs() {
		return true
	}

	hasSameHost := strings.EqualFold(base.Host, target.Host)

	return hasSameHost
}

// IsLinkAccessible checks whether the link is accessible
// Assumption: If the http.Head request timeouts in 5 seconds then the url is inaccessible
func (u *Utils) IsLinkAccessible(link string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Head(link)
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()

	return res.StatusCode < 400
}
