FROM golang:1.14.7-alpine3.12

ENV GO111MODULE=on

ADD ./server /go/src/

EXPOSE 3389

CMD ["/go/src/server"]