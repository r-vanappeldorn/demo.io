FROM golang:1.18-buster

WORKDIR /app

COPY . .

RUN make build

CMD ["./functions-service"]