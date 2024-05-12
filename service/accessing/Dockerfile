FROM golang:1.22

WORKDIR /accessing

COPY . .

RUN go mod tidy

EXPOSE ${CLIENT_PORT}
EXPOSE ${ADMIN_PORT}

RUN mkdir logs && touch ./logs/service_logs.log

VOLUME ./logs
CMD go run ./cmd/main.go
