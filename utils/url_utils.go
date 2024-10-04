package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (u *Utils) FetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("unable to fetch the URL")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("unexpected status code received. status code: %d", resp.StatusCode)
		return "", errors.New(errMsg)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("error reading the response body")
	}

	return string(body), nil
}
