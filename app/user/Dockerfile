# Build stage
FROM golang:1.19
WORKDIR /usr/src

COPY lib/ .
COPY app/user .

RUN go mod download && \
    go mod tidy && \
    go build -o . cmd/main.go && \
    go install github.com/pressly/goose/v3/cmd/goose@latest

ENV GOOSE_DRIVER=postgres
# TODO: Create Authen DB
ENV GOOSE_DBSTRING="host=postgres user=postgres dbname=postgres password=postgres sslmode=disable"

ENTRYPOINT ["/bin/bash","./entrypoint.sh"]

