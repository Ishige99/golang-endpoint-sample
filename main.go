package main

import (
	"log"
	"net/http"
)

func main() {
	// データベースに接続を行います。
	InitDB()

	// サーバーが起動している間はDB接続情報を保持しなければいけないため、main関数内に記述しています。
	defer db.Close()

	// テストエンドポイント
	http.HandleFunc("/", TestHandler)

	// article関連のエンドポイントになります。
	http.HandleFunc("/article", ArticleHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
