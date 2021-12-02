FROM golang:latest

WORKDIR /usr/app

COPY . .

RUN go mod download
RUN go get github.com/cosmtrek/air
