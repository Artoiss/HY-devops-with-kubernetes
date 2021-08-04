FROM golang:alpine

RUN mkdir /app

ADD ./server-go /app

WORKDIR /app

RUN go build -o server

CMD ["/app/server"]