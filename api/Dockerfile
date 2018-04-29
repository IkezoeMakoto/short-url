FROM golang:1.9-alpine

RUN apk update && \
    apk add make git && \
    rm -rf /var/cache/apk/*

RUN apk update && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

COPY ./ /go/src/github.com/IkezoeMakoto/short-url/api

WORKDIR /go/src/github.com/IkezoeMakoto/short-url/api

ARG ENV=local
COPY .env.${ENV} .env

RUN make setup

RUN make build

CMD ["./bin/server"]
