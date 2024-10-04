package main

import (
	"log"
	"net/http"

	"github.com/isurukdniss/webpage-analyzer/handler"
)

func main() {
	fs := http.FileServer(http.Dir("web/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fs))

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/analyze", handler.AnalyzeHandler)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
