package handler

import (
	"net/http"

	"github.com/isurukdniss/webpage-analyzer/analyzer"
	"github.com/isurukdniss/webpage-analyzer/utils"
)

var utilsInstance utils.UtilProvider = &utils.Utils{}
var analyzerInstance analyzer.PageAnalyzer = &analyzer.Analyzer{}
var templatePath = "web/index.html"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := utilsInstance.RenderTemplate(w, r, templatePath, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Execute the analyze logic
		formURL := r.FormValue("url")
		res := analyzerInstance.Analyze(formURL)

		// Render the html
		err := utilsInstance.RenderTemplate(w, r, templatePath, res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}
