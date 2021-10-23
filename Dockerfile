FROM golang:latest

WORKDIR /usr/app

COPY /graphql .
RUN source .env

RUN go mod download
RUN go build -o server server.go

ENV PORT $PORT
EXPOSE $PORT
CMD go run server.go
