# Build stage
FROM golang:1.21
WORKDIR /usr/src

COPY lib/ .
COPY app/order .

RUN go mod download && \
    go mod tidy && \
    go build -o . cmd/main.go 

ENTRYPOINT ["/bin/bash","./entrypoint.sh"]

