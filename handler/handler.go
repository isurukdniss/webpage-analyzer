package handler

import (
	"net/http"

	"github.com/isurukdniss/webpage-analyzer/utils"
)

var templatePath = "web/index.html"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := utils.RenderTemplate(w, r, templatePath, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
}
