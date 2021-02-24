FROM  golang:alpine

COPY . /go/src/github.com/okex/db-traveller

WORKDIR /go/src/github.com/okex/db-traveller

RUN go build -o db-traveller

CMD ["./db-traveller"]



