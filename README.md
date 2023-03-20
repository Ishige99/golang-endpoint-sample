# endpoint-sample

## Env

- MacBook Pro(M1,2021)
- macOS Monterey v12.2.1
- go version go1.20 darwin/arm64
- Docker version 20.10.22
- Docker Compose version v2.15.1
  - mysql  Ver 8.0.32 for Linux on aarch64
  - phpMyAdmin v5.2.1

## Build

- Start Endpoint Server

```shell
% go build .
% ./main
```

- Start MySQL / phpMyAdmin

```shell
% docker-compose up -d
```

## Request

```shell
% curl 'http://localhost:8080/'
// Hello My Server

% curl -X GET 'http://localhost:8080/article' | jq
// JSON Responce (all article)

% curl -X POST -H 'Content-Type: application/json' --data '{"title": "hoge","description": "hogei","content": "hohogehoge"}' 'http://localhost:8080/post-article'
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