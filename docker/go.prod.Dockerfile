FROM golang:latest

WORKDIR /usr/app

COPY /backend .

RUN go mod download
RUN go build -o server server.go

CMD go run server.go
