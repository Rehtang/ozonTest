FROM golang:1.20.5 AS builder

# Установка зависимостей для сборки
RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main /app/main

WORKDIR /app

CMD ["./main"]
