# Stage 1: Build the Go application
FROM golang:1.18 AS build

# Set the working directory in the build stage
WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application, placing the binary in the Server folder
RUN go build -o Server/main ./Server/main.go

# Stage 2: Create a minimal image for the final application
FROM alpine:3.15

# Install necessary libraries for running the Go binary
RUN apk add --no-cache libgcc libstdc++

# Set the working directory
WORKDIR /app/Server

# Copy the Go binary from the build stage
COPY --from=build /app/Server/main .

# Expose the port that the backend listens on
EXPOSE 8080

# Set environment variables if needed
ENV DB_PATH /app/Database
ENV IDENTITY_DB_PATH /app/Database/Identity.db
ENV BUSINESS_DB_PATH /app/Database/Business.db
ENV IDENTITY_MIGRATIONS_PATH /app/Database/migrations/identity
ENV BUSINESS_MIGRATIONS_PATH /app/Database/migrations/business

# Run the application from the Server directory
CMD ["./main"]
