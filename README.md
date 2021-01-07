# go-http-server-sample

a sample HTTP server built with [Echo](https://github.com/labstack/echo)

## Build

``` sh
$ go build
```

### with Docker

``` sh
$ docker build -t go-http-server-sample --build-arg ENV=dev .
```

## Run

``` sh
$ APP_ENV=dev ./go-http-server-sample
```

### with Docker

``` sh
$ docker run -p 1323:1323 -d go-http-server-sample
```
