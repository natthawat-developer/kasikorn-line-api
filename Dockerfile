# Use the official Golang image to build the Go app
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download the dependencies. Dependencies will be cached if the go.mod and go.sum are not changed
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go app
RUN go build -o kasikorn-line-api .

# Start a new stage from the Alpine image
FROM alpine:latest

# Install necessary dependencies for running the Go app (e.g., certificates)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/kasikorn-line-api .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./kasikorn-line-api"]
