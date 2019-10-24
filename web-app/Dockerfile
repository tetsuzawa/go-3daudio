FROM golang:latest

RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work

RUN go get github.com/tetsuzawa/go-3daudio
RUN go get github.com/oklog/ulid
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/satori/go.uuid
RUN go get golang.org/x/crypto/bcrypt

CMD ["go", "run", "main.go"]
