FROM golang:buster

WORKDIR /app
COPY * /app/
RUN go build .