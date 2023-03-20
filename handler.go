package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello My Server")
	fmt.Println("Endpoint Hit: /")
}

// -------受け取ったパラメータに応じて取得す内容を変える記事取得のハンドラ---------

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit '/article', Get all article")

	// articleテーブルから対象のレコードデータを取得します。
	rows, err := db.Query("SELECT id, title, description, content FROM article")

	// データ取得の際にエラーが発生した場合は500エラーを返します。
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 処理終了時に取得したリソースの解放(削除)をします
	// 削除しないとサーバーが起動している限りデータが残り続けます。
	defer rows.Close()

	var articles []Article

	// 先ほどクエリで取得したデータをループでappendしていきます。
	for rows.Next() {
		var article Article
		if err = rows.Scan(&article.ID, &article.Title, &article.Description, &article.Content); err != nil {
			log.Fatal(w, err.Error(), http.StatusInternalServerError)
		}
		articles = append(articles, article)
	}

	// ループ内のエラーを検出し、処理を止めます
	if err = rows.Err(); err != nil {
		log.Fatal(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
