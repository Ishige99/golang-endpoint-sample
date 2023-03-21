package main

// articleをGET、POSTする際に使用する構造体です。

type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
