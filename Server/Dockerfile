# Server/Dockerfile
FROM golang:1.18-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files from the root directory
COPY ../../go.mod ../../go.sum ./

# Download all dependencies
RUN go mod download

# Copy the backend source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal image for the final build
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the frontend build files from the frontend container
COPY --from=frontend /app/build ./build

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
