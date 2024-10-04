package analyzer

import (
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
	mockUtils.EXPECT().IsInternalLink(pageURL, "http://test.com").Return(false, nil).Times(1)
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
