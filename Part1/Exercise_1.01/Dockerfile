FROM golang:alpine

RUN mkdir /app

ADD ./timestamper-go /app

WORKDIR /app


RUN go build -o timestamp .
CMD ["/app/timestamp"]


