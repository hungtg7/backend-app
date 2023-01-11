# Build stage
FROM golang:1.18
WORKDIR /usr/src

COPY . .

RUN make install
RUN make generate
RUN go build -o . app/pet/main.go 

# ENTRYPOINT ["/bin/bash"]
EXPOSE 11000
