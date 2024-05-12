FROM golang:1.22

WORKDIR /task_manager

COPY . .

RUN go mod tidy

EXPOSE ${TASK_PORT}

RUN mkdir logs && touch ./logs/task_logs.log

VOLUME ./logs
CMD go run ./task_manager.go
