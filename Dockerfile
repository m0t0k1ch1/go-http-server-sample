FROM golang:1.15-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o go-http-server-sample

FROM alpine:3.12
ARG APP_ENV=dev
ENV APP_ENV=$APP_ENV
WORKDIR /app
COPY --from=builder /app/go-http-server-sample ./
COPY ./configs/${APP_ENV}.json ./configs/
ENTRYPOINT ["/app/go-http-server-sample"]
