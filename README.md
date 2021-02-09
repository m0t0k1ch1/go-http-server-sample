# go-http-server-sample

[![Go Reference](https://pkg.go.dev/badge/github.com/m0t0k1ch1/go-http-server-sample.svg)](https://pkg.go.dev/github.com/m0t0k1ch1/go-http-server-sample)
[![Documentation Status](https://readthedocs.org/projects/go-http-server-sample/badge/?version=latest)](https://go-http-server-sample.readthedocs.io/en/latest/?badge=latest)

a sample HTTP server built with [Echo](https://github.com/labstack/echo)

ref. [Go で素朴な HTTP API サーバーを書く](https://m0t0k1ch1st0ry.com/blog/2021/01/20/go-http-server-sample)

## Build

```sh
$ docker-compose build
```

## Run

```sh
$ docker-compose up -d
```

## Test

```sh
$ docker-compose exec app go test -v ./...
```
