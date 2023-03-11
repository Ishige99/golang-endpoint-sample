package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", TestHandler)
	http.HandleFunc("/article", GetArticleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
