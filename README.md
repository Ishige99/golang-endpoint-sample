# endpoint-sample

## Env

- go version go1.20 darwin/arm64

## Build & Start Endpoint Server

```shell
% go build .
% ./main
```

## TestRequest

```shell
% curl 'http://localhost:8080/'
// Hello My Server

% curl 'http://localhost:8080/article'
// JSON Responce (article)
```