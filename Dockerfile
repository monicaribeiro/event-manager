FROM golang:1.14-alpine as builder

RUN apk update && apk add gcc libc-dev git

WORKDIR /app/event-manager

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/event-manager .

# -----------------------------------------------

FROM alpine:3.12.0

EXPOSE 8080

ENTRYPOINT ["./event-manager"]