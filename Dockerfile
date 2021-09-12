# FROM golang:latest
# FROM alpine:3.14 as builder
FROM golang:1.17.1-alpine3.14 as builder

# ENV PATH /usr/local/go/bin:$PATH

# ENV GOLANG_VERSION 1.17.1

# airをインストールし、実行権限を与える
RUN apk update \
  && apk add --no-cache git curl \
  && go get -u github.com/cosmtrek/air \
  && chmod +x ${GOPATH}/bin/air

WORKDIR /myapp/backend

# COPY go.mod go.sum ./

# ライブラリのインストール
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /main ./cmd

FROM alpine:3.14

COPY --from=builder /main .

ENV PORT=${PORT}

ENTRYPOINT [ "/main web" ]
