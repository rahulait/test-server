# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy Go module files and source code
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go binary
RUN go build -o server

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/test-server .

# Expose ports (HTTP, TCP, UDP)
EXPOSE 8080 8989 9090
EXPOSE 4343 4545 5656/tcp
EXPOSE 7070 7272 7474/udp

# Run the server
ENTRYPOINT ["./test-server"]
