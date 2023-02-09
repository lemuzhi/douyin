FROM golang:1.18

MAINTAINER "lemuzhi"

WORKDIR /go/src/byte-applets-api

COPY . /go/src/byte-applets-api

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

EXPOSE 8080
ENTRYPOINT ./server
