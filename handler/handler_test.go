package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	gomock "go.uber.org/mock/gomock"

	"github.com/isurukdniss/webpage-analyzer/analyzer"
	analyzerMocks "github.com/isurukdniss/webpage-analyzer/analyzer/mocks"
	"github.com/isurukdniss/webpage-analyzer/utils/mocks"
)

func TestIndexHandler(t *testing.T) {

	// mock the utils package
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUtils := mocks.NewMockUtilProvider(ctrl)

	utilsInstance = mockUtils

	tests := []struct {
		name                   string
		mockRenderTemplateFunc func() error
		expectedStatusCode     int
		expectedBody           string
	}{
		{
			name: "Render Successful",
			mockRenderTemplateFunc: func() error {
				return nil
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       "",
		},
		{
			name: "Error rendering the tempalte",
			mockRenderTemplateFunc: func() error {
				return errors.New("Internal server error")
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       "Internal server error\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockUtils.EXPECT().RenderTemplate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(test.mockRenderTemplateFunc())

			req, _ := http.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(IndexHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != test.expectedStatusCode {
				t.Errorf("Expected status code '%d', got '%d'", test.expectedStatusCode, rr.Code)
			}

			if rr.Body.String() != test.expectedBody {
				t.Errorf("Expected response body '%s', got '%s'", test.expectedBody, rr.Body.String())
			}

		})
	}
}

func TestAnalyzeHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAnalyzer := analyzerMocks.NewMockPageAnalyzer(ctrl)
	mockUtils := mocks.NewMockUtilProvider(ctrl)

	analyzerInstance = mockAnalyzer
	utilsInstance = mockUtils

	formURL := "http://example.com"
	expectedResult := &analyzer.AnalyzerResult{
		HTMLVersion:        "HTML5",
		Title:              "Example Title",
		HeadingsCount:      map[string]int{"h1": 1, "h2": 2},
		InternalLinksCount: 5,
		ExternalLinksCount: 3,
		InAccessibleLinks:  2,
		HasLoginForm:       false,
		ErrorMessage:       "",
		ExternalLinks:      []string{"http://external.com"},
	}

	mockAnalyzer.EXPECT().Analyze(formURL).Return(expectedResult)
	mockUtils.EXPECT().RenderTemplate(gomock.Any(), gomock.Any(), gomock.Any(), expectedResult).Return(nil)

	form := url.Values{"url": {formURL}}
	req := httptest.NewRequest(http.MethodPost, "/analyze", nil)
	req.PostForm = form
	rr := httptest.NewRecorder()

	AnalyzeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code '%d', got '%d'", http.StatusOK, rr.Code)
	}
}
