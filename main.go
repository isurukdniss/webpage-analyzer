package main

import (
	"log"
	"net/http"

	"github.com/isurukdniss/webpage-analyzer/handler"
)

var stylesPathPattern = "/styles/"
var stylesDir = "web/styles"

func main() {
	fs := http.FileServer(http.Dir(stylesDir))
	http.Handle(stylesPathPattern, http.StripPrefix(stylesPathPattern, fs))

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/analyze", handler.AnalyzeHandler)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
