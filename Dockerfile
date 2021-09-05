# Build stage
FROM golang:1.16
WORKDIR /usr/src

COPY . .

RUN make install
RUN make generate
RUN go build -o . app/app_data_monitoring/main.go 

ENV SLACK_WEB_HOOK=$SLACK_WEB_HOOK
# ENTRYPOINT ["/bin/bash"]
EXPOSE 11000