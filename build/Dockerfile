FROM golang:1.14.7-alpine3.12 as builder

WORKDIR /go/src
ENV GO111MODULE=on

RUN mkdir /gobuild
RUN apk add --no-cache alpine-sdk git \
    && git clone https://github.com/YasuhiroOsajima/swaggo-test.git \
    && cd swaggo-test \
    && go build -o /gobuild/server cmd/main.go

# run
FROM alpine:latest

COPY --from=builder /gobuild/server /server
EXPOSE 3389

CMD ["/server"]
