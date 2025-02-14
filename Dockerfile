# Use the official Golang image with a newer version
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Install the latest version of air from the new repository
RUN go install github.com/air-verse/air@latest

# Copy the rest of the application code
COPY . .

# Build the application and place the binary in /app
RUN go build -o /app/sample-health .

# Use a minimal Alpine image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage and make it executable
COPY --from=builder /app/sample-health /app/sample-health
RUN chmod +x /app/sample-health

# Copy the air binary from the builder stage
COPY --from=builder /go/bin/air /usr/local/bin/air

# Copy the air.toml file
COPY air.toml /app/air.toml

# Expose the port the app runs on
EXPOSE 8080

# Run the application with air
CMD ["air", "-c", "air.toml"]