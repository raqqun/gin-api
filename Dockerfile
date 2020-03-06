FROM golang:1.13.7-buster

WORKDIR /go/src/gin-api

COPY . .

RUN go build .

EXPOSE 8080

CMD ["./gin-api"]
