FROM golang:1.18

MAINTAINER "lemuzhi"

WORKDIR /go/src/douyin-api

COPY . /go/src/douyin-api

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
#    && go env \
#    && go mod tidy \
    && go build -o server .

RUN chmod +x ./third_party/ffmpeg.exe

EXPOSE 8080
ENTRYPOINT ./server
