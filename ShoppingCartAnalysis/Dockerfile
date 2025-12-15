# stage 1: Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o shopcart ./cmd/main.go

#stage 2: Runtime stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/shopcart .
COPY .env .

EXPOSE 8080
CMD ["./shopcart"]