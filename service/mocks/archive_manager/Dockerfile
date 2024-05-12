FROM golang:1.22

WORKDIR /archive_manager

COPY . .

RUN go mod tidy

EXPOSE ${ARCHIVE_PORT}

RUN mkdir logs && touch ./logs/archive_logs.log

VOLUME ./logs
CMD go run ./archive_manager.go
