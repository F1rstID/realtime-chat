# Builder stage
FROM golang:1.21-alpine AS builder

# Install required packages and swag CLI
RUN apk add --no-cache gcc musl-dev sqlite-dev git && \
    go install github.com/swaggo/swag/cmd/swag@latest

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies and verify
RUN go mod download && \
    go mod verify

# Copy source code
COPY . .

# Generate Swagger documentation
RUN swag init

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

# Install required runtime packages
RUN apk add --no-cache sqlite-libs ca-certificates tzdata

# Create non-root user
RUN adduser -D -g '' appuser

# Create necessary directories and set permissions
RUN mkdir -p /app/logs /app/docs /app/data && \
    chown -R appuser:appuser /app

# Set working directory
WORKDIR /app

# Copy binary and swagger docs from builder
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Create volume mount points
VOLUME ["/app/data"]

# Set ownership
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose ports
EXPOSE 5050

# Set environment variables
ENV GO_ENV=production
ENV SERVER_PORT=5050
ENV SERVER_URL=0.0.0.0
ENV DATABASE_DSN=/app/data/sqlite.db

# Command to run the application
CMD ["./main"]