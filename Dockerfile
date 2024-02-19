# Build stage
FROM golang:1.22
WORKDIR /usr/src

COPY . .

RUN make install
RUN make generate
RUN go build -o . app/app_data_monitoring/main.go 

# ENTRYPOINT ["/bin/bash"]
EXPOSE 11000
