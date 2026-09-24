FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o notification-agent ./cmd/server

CMD ["./notification-agent"]