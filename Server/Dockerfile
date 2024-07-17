# Use official Golang image with Alpine Linux as base
FROM golang:1.18-alpine AS build

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files from the root directory and download dependencies
COPY ../go.mod ../go.sum ./
RUN go mod download

# Copy the source code from the Server directory into the container
COPY ../ .

# Move to the Server directory
WORKDIR /app/Server

# Build the Go app
# RUN go build -o main .
RUN go build -o /app/main .

# Stage 2: Use Alpine Linux as the base for the runtime environment
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/main .

# Copy the Database directory to the container
COPY --from=build /app/Database /app/Database

# Expose port 8080 to outside traffic
EXPOSE 8080

# Command to run the backend server
CMD ["./main"]
