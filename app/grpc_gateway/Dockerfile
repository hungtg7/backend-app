# Build stage
FROM golang:1.21
WORKDIR /usr/src

COPY lib/ .
COPY app/grpc_gateway .

RUN go mod download && \
    go mod tidy && \
    go build -o . ./main.go 

ENTRYPOINT ["/bin/bash","./entrypoint.sh"]

