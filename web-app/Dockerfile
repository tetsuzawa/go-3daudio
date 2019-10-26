FROM golang:latest

RUN mkdir -p /go/src/github.com/tetsuzawa/go-3daudio/web-app
WORKDIR /go/src/github.com/tetsuzawa/go-3daudio/web-app
ADD . /go/src/github.com/tetsuzawa/go-3daudio/web-app

RUN go get github.com/oklog/ulid
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/satori/go.uuid
RUN go get golang.org/x/crypto/bcrypt
RUN go get gopkg.in/ini.v1
RUN go get github.com/joho/godotenv

CMD ["go", "run", "main.go"]
