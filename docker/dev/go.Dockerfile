FROM golang:latest

WORKDIR /usr/app

COPY . .

RUN go mod download
RUN go install github.com/cosmtrek/air
