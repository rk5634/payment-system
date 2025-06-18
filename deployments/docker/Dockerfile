# Build stage
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /payment ./cmd/paymentservice

# Final image
FROM alpine:latest

RUN adduser -D -g '' payment
USER payment

WORKDIR /home/payment

COPY --from=builder /payment .

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget --spider -q http://localhost:8080/health || exit 1

CMD ["./payment"]
