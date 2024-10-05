// Code generated by MockGen. DO NOT EDIT.
// Source: utils/util_provider.go
//
// Generated by this command:
//
//	mockgen -source=utils/util_provider.go -destination=utils/mocks/mock_util_provider.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	html "golang.org/x/net/html"
)

// MockUtilProvider is a mock of UtilProvider interface.
type MockUtilProvider struct {
	ctrl     *gomock.Controller
	recorder *MockUtilProviderMockRecorder
}

// MockUtilProviderMockRecorder is the mock recorder for MockUtilProvider.
type MockUtilProviderMockRecorder struct {
	mock *MockUtilProvider
}

// NewMockUtilProvider creates a new mock instance.
func NewMockUtilProvider(ctrl *gomock.Controller) *MockUtilProvider {
	mock := &MockUtilProvider{ctrl: ctrl}
	mock.recorder = &MockUtilProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUtilProvider) EXPECT() *MockUtilProviderMockRecorder {
	return m.recorder
}

// ExtractAttribute mocks base method.
func (m *MockUtilProvider) ExtractAttribute(n *html.Node, attr string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractAttribute", n, attr)
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtractAttribute indicates an expected call of ExtractAttribute.
func (mr *MockUtilProviderMockRecorder) ExtractAttribute(n, attr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractAttribute", reflect.TypeOf((*MockUtilProvider)(nil).ExtractAttribute), n, attr)
}

// ExtractHTMLVersion mocks base method.
func (m *MockUtilProvider) ExtractHTMLVersion(htmlContent string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractHTMLVersion", htmlContent)
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtractHTMLVersion indicates an expected call of ExtractHTMLVersion.
func (mr *MockUtilProviderMockRecorder) ExtractHTMLVersion(htmlContent any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractHTMLVersion", reflect.TypeOf((*MockUtilProvider)(nil).ExtractHTMLVersion), htmlContent)
}

// ExtractTitle mocks base method.
func (m *MockUtilProvider) ExtractTitle(n *html.Node) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractTitle", n)
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtractTitle indicates an expected call of ExtractTitle.
func (mr *MockUtilProviderMockRecorder) ExtractTitle(n any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractTitle", reflect.TypeOf((*MockUtilProvider)(nil).ExtractTitle), n)
}

// FetchURL mocks base method.
func (m *MockUtilProvider) FetchURL(url string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchURL", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchURL indicates an expected call of FetchURL.
func (mr *MockUtilProviderMockRecorder) FetchURL(url any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchURL", reflect.TypeOf((*MockUtilProvider)(nil).FetchURL), url)
}

// HasLoginForm mocks base method.
func (m *MockUtilProvider) HasLoginForm(n *html.Node) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasLoginForm", n)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasLoginForm indicates an expected call of HasLoginForm.
func (mr *MockUtilProviderMockRecorder) HasLoginForm(n any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasLoginForm", reflect.TypeOf((*MockUtilProvider)(nil).HasLoginForm), n)
}

// IsInternalLink mocks base method.
func (m *MockUtilProvider) IsInternalLink(baseURL, targetURL string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInternalLink", baseURL, targetURL)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInternalLink indicates an expected call of IsInternalLink.
func (mr *MockUtilProviderMockRecorder) IsInternalLink(baseURL, targetURL any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInternalLink", reflect.TypeOf((*MockUtilProvider)(nil).IsInternalLink), baseURL, targetURL)
}

// IsLinkAccessible mocks base method.
func (m *MockUtilProvider) IsLinkAccessible(link string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLinkAccessible", link)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLinkAccessible indicates an expected call of IsLinkAccessible.
func (mr *MockUtilProviderMockRecorder) IsLinkAccessible(link any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLinkAccessible", reflect.TypeOf((*MockUtilProvider)(nil).IsLinkAccessible), link)
}

// ParseHTML mocks base method.
func (m *MockUtilProvider) ParseHTML(pageHTML string) (*html.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseHTML", pageHTML)
	ret0, _ := ret[0].(*html.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseHTML indicates an expected call of ParseHTML.
func (mr *MockUtilProviderMockRecorder) ParseHTML(pageHTML any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseHTML", reflect.TypeOf((*MockUtilProvider)(nil).ParseHTML), pageHTML)
}

// RenderTemplate mocks base method.
func (m *MockUtilProvider) RenderTemplate(w http.ResponseWriter, r *http.Request, templatePath string, data any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderTemplate", w, r, templatePath, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenderTemplate indicates an expected call of RenderTemplate.
func (mr *MockUtilProviderMockRecorder) RenderTemplate(w, r, templatePath, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderTemplate", reflect.TypeOf((*MockUtilProvider)(nil).RenderTemplate), w, r, templatePath, data)
}
