# Build stage
FROM golang:1.19
WORKDIR /usr/src

COPY lib/ .
COPY app/pet .

RUN go mod download
RUN go build -o . cmd/main.go 

ENTRYPOINT ["/bin/bash","./entrypoint.sh"]

