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

EXPOSE 8080
ENTRYPOINT ./server

#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64 \
#        GOPROXY="https://goproxy.cn,direct"
#
#MAINTAINER "lemuzhi"
#
#WORKDIR /go/src/byte-applets-api
#
#COPY . .
#RUN go get -d -v ./...
#RUN go install -v ./...
#
#EXPOSE 12345:12345
#CMD ["byte-applets-api"]
