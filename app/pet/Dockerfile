# Build stage
FROM golang:1.18
WORKDIR /usr/src

COPY lib/ .
COPY app/pet .
COPY Makefile .

# RUN make generate
ENV GOROOT=/usr/src
RUN go build -o . cmd/main.go 

# ENTRYPOINT ["/bin/bash"]
