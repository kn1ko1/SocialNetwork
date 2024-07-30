# Stage 1: Build the Go application
FROM golang:1.18 AS build

# Install build essentials
RUN apt-get update && apt-get install -y gcc libc6-dev

# Set the working directory
WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application with CGO enabled
RUN CGO_ENABLED=1 go build -o /app/Server/main ./Server/main.go

# Stage 2: Create a minimal image for the final application
FROM alpine:3.15

# Install necessary libraries
RUN apk add --no-cache libgcc libstdc++

# Set the working directory
WORKDIR /app/Server

# Copy the Go binary from the build stage
COPY --from=build /app/Server/main .

# Expose the port that the backend listens on
EXPOSE 8080

# Set environment variables
ENV DB_PATH=/app/Database \
    IDENTITY_DB_PATH=/app/Database/Identity.db \
    BUSINESS_DB_PATH=/app/Database/Business.db \
    IDENTITY_MIGRATIONS_PATH=/app/Database/migrations/identity \
    BUSINESS_MIGRATIONS_PATH=/app/Database/migrations/business

# Run the application
CMD ["./main"]