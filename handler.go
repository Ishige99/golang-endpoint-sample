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

// -------------------------------------
// -------メソッドごとに分岐をします---------
// -------------------------------------

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("Endpoint Hit '/article', GET All Articles")
		GetArticleHandler(w, r)
	case http.MethodPost:
		fmt.Println("Endpoint Hit '/article', POST Article")
		CreateArticleHandler(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed) // 405
	}
}

// -----------------------------
// -------記事投稿ハンドラ---------
// -----------------------------

func CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	// POSTパラメータのJSON値をGoで扱える状態(JSONデコーダ)に変換
	decoder := json.NewDecoder(r.Body)

	// リクエストボディも保持し続けるとメモリを圧迫するので、ハンドラ関数の終了時にデータは解放をする
	defer r.Body.Close()

	var newArticle Article

	// JSONデコーダからnewArticle構造体にデコードを行います
	err := decoder.Decode(&newArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// クエリを実行して、articleテーブルにINSERTします
	result, err := db.Exec(
		`INSERT INTO article (title, description, content)
		VALUES 	(?, ?, ?);`,
		newArticle.Title,
		newArticle.Description,
		newArticle.Content,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスに出力するため、最後に投稿したIDを取得
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Article created with ID: %d\n", lastInsertedID)
}

// -----------------------------
// -------記事取得ハンドラ---------
// -----------------------------

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
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
