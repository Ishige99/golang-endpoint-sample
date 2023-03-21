package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// TODO: 部分的にlog.Fatal()を使用してますが、これだとエラー発生時に処理が終了してしまうので、log.Printlnやlog.Printfを使用してもいいかも

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello My Server")
	fmt.Println("Endpoint Hit: /")
}

// -------------------------------------
// -------メソッドごとに分岐をします---------
// -------------------------------------

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストパラメータに`id`があるか確認をします。
	hasID := r.URL.Query().Has("id")

	// リクエストの条件によって、ハンドラを切り替えます。
	switch r.Method {
	case http.MethodGet:
		// パラメータに`id`があれば単一記事取得
		if hasID {
			fmt.Println("Endpoint Hit '/article', GET Single Article")
			GetSingleArticleHandler(w, r)
			return
		}
		// それ以外であれば全体記事取得
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

	// リクエストボディも保持し続けるとメモリを圧迫するので、ハンドラ関数の終了時にデータは解放をします
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

// --------------------------------
// -------全体記事取得ハンドラ---------
// --------------------------------

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	// articleテーブルから対象のレコードデータを取得します。
	rows, err := db.Query("SELECT id, title, description, content FROM article")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	// 処理終了時に取得したリソースの解放(削除)をします。※削除しないとサーバーが起動している限りデータが残り続けます。
	defer rows.Close()

	var articles []Article

	// 先ほどクエリで取得したデータをループでappendしていきます。
	for rows.Next() {
		var article Article
		if err = rows.Scan(&article.ID, &article.Title, &article.Description, &article.Content); err != nil {
			log.Fatal(w, err.Error(), http.StatusInternalServerError) // 500
		}
		articles = append(articles, article)
	}

	// ループ内のエラーを検出した場合は、この時点で処理を終了します。
	if err = rows.Err(); err != nil {
		log.Fatal(w, err.Error(), http.StatusInternalServerError) // 500
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// --------------------------------
// -------単一記事取得ハンドラ---------
// --------------------------------

func GetSingleArticleHandler(w http.ResponseWriter, r *http.Request) {
	// パラメータから`id`の値を取得し、文字列から数値に変換します。
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)

	// 文字列, 0以下のパラメータが含まれていた場合にエラーを返します。
	if err != nil || id <= 0 {
		http.Error(w, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// パラメータから取得した`id`の情報をもとに対象の記事取得クエリを発行、データ抽出を行います
	var row *sql.Rows
	row, err = db.Query("SELECT id, title, description, content FROM article WHERE id = ?", id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	// 処理終了時に取得したリソースの解放(削除)をします。※削除しないとサーバーが起動している限りデータが残り続けます。
	defer row.Close()

	var article Article

	// クエリで取得したデータをスキャンしてレスポンスで返す変数に入れていきます。
	if row.Next() {
		if err = row.Scan(&article.ID, &article.Title, &article.Description, &article.Content); err != nil {
			log.Fatal(w, err.Error(), http.StatusInternalServerError) // 500
		}
	} else {
		http.Error(w, "Article not found", http.StatusNotFound) // 404
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}
