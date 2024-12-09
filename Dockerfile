# Build Stage
FROM golang:1.22.10 AS builder

WORKDIR /app
COPY . /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the project files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/loan/main.go
RUN ls -l /app

# Runtime Stage
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x /root/main

# Expose the port
EXPOSE 9090

# Run the application
CMD ["./main"]
