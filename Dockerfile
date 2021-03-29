FROM golang:1.16-alpine as builder

RUN mkdir /build

COPY . /build/go-placeholder-client-www

RUN apk update && apk upgrade \
    && apk add make git \
    && cd /build/go-placeholder-client-www \
    && go build -mod vendor -o /usr/local/bin/placeholder-client-www cmd/server/main.go    

FROM alpine:latest

COPY --from=builder /usr/local/bin/placeholder-client-www /usr/local/bin/

RUN apk update && apk upgrade \
    && apk add ca-certificates