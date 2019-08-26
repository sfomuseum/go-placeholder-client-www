FROM golang:1.12-alpine as builder

RUN mkdir /build

COPY . /build/go-placeholder-client-www

RUN apk update && apk upgrade \
    && apk add make git \
    && cd /build/go-placeholder-client-www \
    && go build -o /usr/local/bin/placeholder-client-www cmd/server/main.go    

FROM alpine:latest

COPY --from=builder /usr/local/bin/placeholder-client-www /usr/local/bin/

RUN apk update && apk upgrade \
    && apk add ca-certificates