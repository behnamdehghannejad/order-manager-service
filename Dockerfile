FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o order-managenment-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/order-managenment-service .

EXPOSE 8080

CMD ["./order-managenment-service"]