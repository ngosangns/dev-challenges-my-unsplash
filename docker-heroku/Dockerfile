FROM golang:alpine

RUN mkdir /go/src/app
COPY . /go/src/app

WORKDIR /go/src/app

CMD go build -o app; ./app