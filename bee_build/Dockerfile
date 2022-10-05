FROM golang:1.14-alpine
RUN apk add --no-cache git openssh-client build-base
ARG GO111MODULE=on
RUN set -x && go get\
  github.com/astaxie/beego@v1.12.3\
  github.com/beego/bee@v1.12.3\
  && bee version
WORKDIR /go/src
