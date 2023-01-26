# Build stage
FROM golang:1.19
WORKDIR /usr/src

COPY lib/ .
COPY app/order .
COPY Makefile .

ENTRYPOINT ["/bin/bash","./entrypoint.sh"]

