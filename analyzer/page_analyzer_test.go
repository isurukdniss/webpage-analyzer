package analyzer

import (
	"errors"
	"strings"
	"testing"

	gomock "go.uber.org/mock/gomock"
	"golang.org/x/net/html"

	"github.com/isurukdniss/webpage-analyzer/utils/mocks"
)

var pageAnalyzer PageAnalyzer = &Analyzer{}

func TestAnalyze(t *testing.T) {
	// mock the utils package
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUtils := mocks.NewMockUtilProvider(ctrl)
	utilsInstance = mockUtils

	pageURL := "http://example.com"
	body := `"<html>
				<head>
					<title>Test Page</title>
				</head>
				<body>
					<a href="http://test.com">Example Link</a>
				</body>
			</html>"`

	doc, _ := html.Parse(strings.NewReader(body))

	expectedTitle := "Test Title"
	expectedHTMLVersion := "HTML 5"
	expectedInternalLinksCount := 0
	expectedExternalLinksCount := 1
	expectedInaccessibleLinksCount := 0

	// setup mocks
	mockUtils.EXPECT().FetchURL(pageURL).Return(body, nil)
	mockUtils.EXPECT().ParseHTML(body).Return(doc, nil)
	mockUtils.EXPECT().ExtractHTMLVersion(body).Return(expectedHTMLVersion)
	mockUtils.EXPECT().ExtractTitle(gomock.Any()).Return(expectedTitle).Times(1)
	mockUtils.EXPECT().ExtractAttribute(gomock.Any(), "href").Return("http://test.com").Times(1)
	mockUtils.EXPECT().IsInternalLink(pageURL, "http://test.com").Return(false).Times(1)
	mockUtils.EXPECT().IsLinkAccessible(gomock.Any()).Return(true)

	res := pageAnalyzer.Analyze(pageURL)

	if res.Title != expectedTitle {
		t.Errorf("Expected title %s, got %s", expectedTitle, res.Title)
	}
	if res.HTMLVersion != expectedHTMLVersion {
		t.Errorf("Expected HTML version %s, got %s", expectedHTMLVersion, res.HTMLVersion)
	}
	if res.InternalLinksCount != expectedInternalLinksCount {
		t.Errorf("Expected internal links count %d, got %d", expectedInternalLinksCount, res.InternalLinksCount)
	}
	if res.ExternalLinksCount != expectedExternalLinksCount {
		t.Errorf("Expected external links count %d, got %d", expectedExternalLinksCount, res.ExternalLinksCount)
	}
	if res.InternalLinksCount != expectedInaccessibleLinksCount {
		t.Errorf("Expected inaccessible links count %d, got %d", expectedInaccessibleLinksCount, res.InternalLinksCount)
	}

}

func TestHandleErrorMsg(t *testing.T) {
	tests := []struct {
		err      error
		expected string
	}{
		{
			err:      errors.New("invalid URI for request"),
			expected: "The provided URL is not valid. Please check the format and try again.",
		},
		{
			err:      errors.New("invalid URL: missing scheme or host"),
			expected: "The URL is missing a scheme (like 'http' or 'https') or a host. Please provide a complete URL.",
		},
		{
			err:      errors.New("unable to fetch the URL"),
			expected: "We were unable to fetch the requested URL. Please check your internet connection or the URL.",
		},
		{
			err:      errors.New("server returned status code 404"),
			expected: "The server returned a status code of 404. Please ensure you have the necessary permissions.",
		},
		{
			err:      errors.New("error reading the response body"),
			expected: "An error occurred while reading the response. Please try again later.",
		},
		{
			err:      errors.New("some unexpected error"),
			expected: "An unexpected error occurred. Please try again.",
		},
	}

	for _, test := range tests {
		t.Run(test.err.Error(), func(t *testing.T) {
			result := handleErrorMsg(test.err)
			if result != test.expected {
				t.Errorf("Expected error message %q, got %q", test.expected, result)
			}
		})
	}
}
