FROM golang:1.15-alpine as build
WORKDIR /app
COPY . .
RUN go build -o app main.go

FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/app /app/app
ENTRYPOINT ["/app/app"]
