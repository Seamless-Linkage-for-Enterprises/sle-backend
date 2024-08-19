# Use the official Golang image as the base image
FROM golang:1.22.5-alpine3.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy all the project files into the container
COPY . .

# Download all dependencies
RUN go mod download

# Build the Go app
RUN go build -o main ./cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main /main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/main"]
