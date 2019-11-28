FROM golang:latest

RUN mkdir -p /go/src/go-3daudio/web-app
WORKDIR /go/src/go-3daudio/web-app
ADD . /go/src/go-3daudio/web-app

RUN go build

CMD ["./web-app"]
