# FROM golang:latest
# FROM alpine:3.14 as builder
FROM golang:1.17.1-alpine3.14 as builder

# ENV PATH /usr/local/go/bin:$PATH

# ENV GOLANG_VERSION 1.17.1

# airをインストールし、実行権限を与える
RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/cosmtrek/air \
  && chmod +x ${GOPATH}/bin/air

WORKDIR /myapp/backend

# COPY go.mod go.sum ./

# ライブラリのインストール
RUN go mod download

COPY . .
