# endpoint-sample

## Development Environment

- MacBook Pro(M1,2021)
- macOS Monterey v12.2.1
- go version go1.20 darwin/arm64
- Docker version 20.10.22
- Docker Compose version v2.15.1
  - mysql  Ver 8.0.32 for Linux on aarch64
  - phpMyAdmin v5.2.1

## Directory Structure

```shell
.
├── README.md
├── db
│   └── sql
│       ├── create_database.sql
│       └── create_tables.sql
├── db.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── handler.go
├── main.go
└── struct.go 
```

## Build

- Start DB Server: MySQL/phpMyAdmin

```shell
% docker-compose up -d
// MySQL: 'localhost:3307'
// phpMyAdmin: 'localhost:4040'

% docker exec -i -t endpoint-sample-mysql /bin/bash
% mysql -u root -p
Enter password: root
// Connect MySQL
```

- Start API Server

```shell
% go build .
% ./main
```

## Request

```shell
% curl 'http://localhost:8080/'
// Hello My Server

% curl -X GET 'http://localhost:8080/article' | jq
// GET all articles

% curl -X GET 'http://localhost:8080/article?id=1' | jq
// GET single articles

% curl -X POST 'http://localhost:8080/article' -H 'Content-Type: application/json' --data '{"title": "hoge","description": "hohoge","content": "hogehoge"}'
// POST article data
```

## Reference

- [Goで簡易的なAPIサーバーを立てる](https://qiita.com/entaku0818/items/c29add790718c215381e)
- [[初心者向け] Golang でシンプルな JSON API を作る](https://zenn.dev/tatsurom/articles/golang-simple-json-api)
- [【Docker】MySQLを簡単に構築](https://zenn.dev/re24_1986/articles/153cdc5db96dc0)
- [【Docker】MySQLワークベンチに繋がらない…DBコンテナのポートフォワーディングの設定](https://qiita.com/ryuji-oda/items/c3ed1b86fe0c1f2b9058)
- [Go言語でデータベース（MySQL）に接続する方法](https://nishinatoshiharu.com/connect-go-database/)
- [database/sqlによるデータベース操作](https://www.wakuwakubank.com/posts/869-go-database-sql/#index_id5)
- [phpMyAdmin / Docs > インストール](https://docs.phpmyadmin.net/ja/latest/setup.html)
- [sql.Openを別の関数に切り出したら「sql: database is closed」エラー](https://qiita.com/obr-note/items/7e4cb141a86cb7c58388)
- [SHOW CREATE TABLE したら panic: sql: expected 4 destination arguments in Scan, not 2 と言われる](https://qiita.com/yuya_takeyama/items/ea688bba37d935e1e510)
- [Goのnet/httpパッケージを使ったAPI開発を素振りしてた](https://yuzu441.hateblo.jp/entry/2019/01/31/232445)