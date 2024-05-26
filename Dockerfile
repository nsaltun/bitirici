FROM golang:1.22.2-bullseye

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy