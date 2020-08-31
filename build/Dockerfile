FROM golang:1.14.7-alpine3.12 as builder

WORKDIR /go/src
ENV GO111MODULE=on

RUN mkdir /gobuild
RUN apk add --no-cache alpine-sdk git \
    && git clone https://github.com/YasuhiroOsajima/swaggo-test.git \
    && cd swaggo-test \
    && go build -o build/server cmd/main.go \
    && go get -u github.com/swaggo/swag/cmd/swag \
    && swag init -g cmd/main.go

# run
FROM alpine:latest

COPY --from=builder /go/src/swaggo-test /swaggo-test/
EXPOSE 3389

WORKDIR /swaggo-test
CMD ["build/server"]
