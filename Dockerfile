# Use the official Golang image as a base image
FROM golang:latest
# Set the working directory inside the container
WORKDIR /app
# Copy only the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
# Copy the local code to the container
COPY . .
# Build the Go application
RUN go build -o social-network .
# Expose the port the app runs on
EXPOSE 8080
# Command to run the executable
CMD ["./social-network"]