# endpoint-sample

## Env

- MacBook Pro(M1,2021)
- macOS Monterey v12.2.1
- go version go1.20 darwin/arm64
- Docker version 20.10.22

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

## TestRequest

```shell
% curl 'http://localhost:8080/'
// Hello My Server

% curl 'http://localhost:8080/article'
// JSON Responce (article)
```

## Reference

- [Goで簡易的なAPIサーバーを立てる](https://qiita.com/entaku0818/items/c29add790718c215381e)
- [【Docker】MySQLを簡単に構築](https://zenn.dev/re24_1986/articles/153cdc5db96dc0)
- [【Docker】MySQLワークベンチに繋がらない…DBコンテナのポートフォワーディングの設定](https://qiita.com/ryuji-oda/items/c3ed1b86fe0c1f2b9058)
- [Go言語でデータベース（MySQL）に接続する方法](https://nishinatoshiharu.com/connect-go-database/)
- [phpMyAdmin / Docs > インストール](https://docs.phpmyadmin.net/ja/latest/setup.html)