FROM golang:1.26

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o notification-agent ./cmd/server

EXPOSE 50051

CMD ["./notification-agent"]