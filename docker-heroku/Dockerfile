FROM golang:alpine

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY . .

ENTRYPOINT go build -o app; ./app