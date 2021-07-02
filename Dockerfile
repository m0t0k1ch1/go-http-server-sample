FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .
RUN apk update && apk add --no-cache bash gcc musl-dev tzdata
RUN go build -o go-http-server-sample
RUN go get -u github.com/cosmtrek/air

FROM alpine:3.14
WORKDIR /app
RUN apk update && apk add --no-cache tzdata
COPY --from=builder /app/go-http-server-sample ./
ENTRYPOINT ["/app/go-http-server-sample"]
