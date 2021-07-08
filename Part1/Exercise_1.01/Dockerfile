FROM golang:1.16

RUN mkdir /app

ADD ./timestamper-go /app

WORKDIR /app


RUN go build -o timestamp .
CMD ["/app/timestamp"]


