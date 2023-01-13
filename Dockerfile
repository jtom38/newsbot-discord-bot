FROM golang:1.19 as build

COPY . /app
WORKDIR /app
RUN go build .

FROM alpine:latest as app

CMD [ "/app/newsbot-discord-bot" ]