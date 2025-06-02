FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o order-service ./cmd/order
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/order-service .
EXPOSE 8080
CMD ["./order-service"]