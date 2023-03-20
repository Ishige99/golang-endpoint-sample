package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() {
	// DSNの作成(DB接続情報)
	dsn := "root:root@tcp(127.0.0.1:3307)/test_db"

	// MySQLに接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err) // エラーがあればログに出力して終了
	}

	// 関数終了時に、DBへの接続を終了する。
	// main関数で使用されているので、main関数の終了次第DBの接続を解除する
	defer db.Close()

	// 実際にDBとの接続を確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err) // エラーがあればログに出力して終了
	} else {
		log.Println("Success database connect.")
	}
}
