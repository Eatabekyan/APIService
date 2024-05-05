FROM golang:1.22 AS build

WORKDIR /accessing

COPY . .

RUN go mod tidy

EXPOSE ${CLIENT_PORT}
EXPOSE ${ADMIN_PORT}


CMD go run ./main/main.go
