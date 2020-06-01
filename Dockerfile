FROM golang:1.14.3-alpine3.11

RUN apk add --no-cache git gcc musl-dev

RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega
RUN go get github.com/jinzhu/copier

RUN mkdir -p /app
WORKDIR /app

ADD . /app
