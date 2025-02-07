# 🏗 Stage 1: Build the Go application
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/kube-event-logger main.go

# 🏗 Stage 2: Create a lightweight final image
FROM debian:latest

# Install PostgreSQL client tools and CA certificates
RUN apt-get update && apt-get install -y \
    postgresql-client \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /root/

# Copy the built binary and migrations
COPY --from=builder /app/kube-event-logger .
COPY --from=builder /app/migrations ./migrations

# Ensure the binary is executable
RUN chmod +x /root/kube-event-logger

# Expose port 8080 for the API
EXPOSE 8080

# Run the application
CMD ["/root/kube-event-logger"]