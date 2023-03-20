package main

import (
	"log"
	"net/http"
)

func main() {
	// データベースに接続を行います。
	InitDB()

	// DB接続情報はmain関数の実行後に終了します。（サーバーが起動している間は保持し続けます）
	defer db.Close()

	http.HandleFunc("/", TestHandler)

	// articleテーブルの全カラムを取り出すエンドポイント
	http.HandleFunc("/article", GetArticleHandler)

	// articleテーブルにPOSTするエンドポイント
	http.HandleFunc("/post-article", CreateArticleHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
