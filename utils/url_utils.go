package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (u *Utils) FetchURL(rawUrl string) (string, error) {
	parsedUrl, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return "", err
	}

	if parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		return "", errors.New("invalid URL: missing scheme or host")
	}

	resp, err := http.Get(rawUrl)
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
