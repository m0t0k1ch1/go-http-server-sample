FROM golang:1.15-alpine as builder
ARG APP_ENV=dev
ENV APP_ENV=$APP_ENV
WORKDIR /app
COPY . .
RUN apk update && apk add --no-cache bash gcc musl-dev tzdata
RUN go build -o go-http-server-sample
RUN go get -u github.com/cosmtrek/air

FROM alpine:3.12
ARG APP_ENV=dev
ENV APP_ENV=$APP_ENV
WORKDIR /app
RUN apk update && apk add --no-cache tzdata
COPY --from=builder /app/go-http-server-sample ./
COPY --from=builder /app/configs/${APP_ENV}.json ./configs/
ENTRYPOINT ["/app/go-http-server-sample"]
