package main

import (
	"log"
	"net/http"

	"github.com/isurukdniss/webpage-analyzer/handler"
)

func main() {
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/analyzer", handler.AnalyzeHandler)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
