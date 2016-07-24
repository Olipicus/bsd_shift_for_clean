FROM golang:alpine

RUN apk update && apk add git && rm -rf /var/cache/apk/*

ADD . /go/src/code.olipicus.com/bsd_shift_for_clean/api
WORKDIR /go/src/code.olipicus.com/bsd_shift_for_clean/api

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/websocket
RUN go get gopkg.in/mgo.v2
RUN go get git.apache.org/thrift.git/lib/go/thrift
RUN go build -o api .
CMD ["./api","-state=production"]
