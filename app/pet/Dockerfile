# Build stage
FROM golang:1.19
WORKDIR /usr/src

COPY lib/ .
COPY app/pet .

RUN go mod download && \
    go build -o . cmd/main.go && \
    go install github.com/pressly/goose/v3/cmd/goose@latest 

ENV GOOSE_DRIVER=postgres
ENV GOOSE_DBSTRING="user=postgres dbname=postgres sslmode=disable"

ENTRYPOINT ["/bin/bash","./entrypoint.sh"]

