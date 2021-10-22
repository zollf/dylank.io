FROM golang:latest

WORKDIR /usr/app

COPY /graphql .

RUN go mod download
RUN go build -o server server.go

CMD go run server.go
