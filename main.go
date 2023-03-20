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

	// テストエンドポイント
	http.HandleFunc("/", TestHandler)

	// article関連のエンドポイントになります。
	http.HandleFunc("/article", ArticleHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
