FROM golang:1.21-alpine

RUN mkdir /provider
ADD . /provider/
WORKDIR /provider

ENV PATH="${PATH}:/app"

RUN go mod download && go mod verify
RUN go build -o provider main.go

CMD ["./provider"]