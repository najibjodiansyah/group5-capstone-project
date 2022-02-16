FROM golang:1.17

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go build -o api

CMD ./api