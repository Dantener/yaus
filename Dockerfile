FROM golang:1.18

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main ./cmd/main.go

ENTRYPOINT ["./main"]