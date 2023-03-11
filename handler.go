package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello My Server")
	fmt.Println("Endpoint Hit: /")
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	var articles []Article
	for i := 0; i < 10; i++ {
		title := "id: %d"
		articles = append(
			articles,
			Article{Title: fmt.Sprintf(title, i), Desc: "Article Description", Content: "Article Content"})
	}

	fmt.Println("Endpoint Hit: /article")
	json.NewEncoder(w).Encode(articles)
}
