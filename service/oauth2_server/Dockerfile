FROM golang:1.22 AS build

WORKDIR /go/src/oauth2_server

COPY . .

RUN go mod tidy

EXPOSE ${OAUTH_SERVER_PORT}

CMD go run ./server.go
