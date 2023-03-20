package main

import (
	"log"
	"net/http"
)

func main() {
	// データベースに接続を行う
	// 成功：Success database connect.
	// 失敗：エラー出力しサーバー自体立ち上がりません。
	InitDB()

	http.HandleFunc("/", TestHandler)

	// articleテーブルの全カラムを取り出すエンドポイント
	// 'http://localhost:8080/article'
	http.HandleFunc("/article", GetArticleHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
