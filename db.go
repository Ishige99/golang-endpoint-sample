package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3307)/test_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// DBに接続
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Success database connect.")
	}
}
