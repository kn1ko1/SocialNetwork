# Stage 1: Build the Go application
FROM golang:1.18 AS build

# Install build essentials
RUN apt-get update && apt-get install -y gcc libc6-dev

# Set the working directory in the build stage
WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application with CGO enabled, placing the binary in the Server directory
RUN CGO_ENABLED=1 go build -o /app/Server/main ./Server/main.go

# Debugging: List files in /app and /app/Server
RUN ls -la /app
RUN ls -la /app/Server

# Stage 2: Create a minimal image for the final application
FROM debian:bullseye-slim

# Set the working directory in the final stage
WORKDIR /app

# Copy the Go binary from the build stage
COPY --from=build /app/Server/main /app/Server/main

# Expose the port that the backend listens on
EXPOSE 8080

# Set environment variables
ENV DB_PATH=/app/Database \
    IDENTITY_DB_PATH=/app/Database/Identity.db \
    BUSINESS_DB_PATH=/app/Database/Business.db \
    IDENTITY_MIGRATIONS_PATH=/app/Database/migrations/identity \
    BUSINESS_MIGRATIONS_PATH=/app/Database/migrations/business

# Ensure the main binary has executable permissions
RUN chmod 755 /app/Server/main

WORKDIR /app/Server

# Run the application
CMD ["./main"]
