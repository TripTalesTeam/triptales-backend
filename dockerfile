# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main ./cmd

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy binary
COPY --from=builder /app/main .

# ✅ Copy the .env file (assuming it's in the project root)
COPY .env .env

EXPOSE 8080

CMD ["./main"]
