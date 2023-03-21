package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB接続情報をグローバル変数にすることで、handler.goなどの他ファイルでも呼び出せるようにしてます
var db *sql.DB

func InitDB() {
	// DSNの作成(DB接続情報)
	// TODO: 環境変数で持ってもいいかも
	dsn := "root:root@tcp(127.0.0.1:3307)/test_db"

	var err error

	// MySQLに接続
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err) // エラーがあればログに出力して終了
	}

	// 実際にDBとの接続を確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err) // エラーがあればログに出力して終了
	} else {
		log.Println("Success database connect.")
	}
}
