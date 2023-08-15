# Start with a minimal base image
FROM golang:latest as builder

# Install Git
RUN apt-get update && apt-get install -y git

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -o myapp

# Create a new image
FROM alpine:latest

# Install Git
RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Copy the config file into the image
COPY . .

# Expose the necessary port(s)
EXPOSE 8181

# Set any necessary environment variables
ENV ENVIRONMENT=local

# Run the binary
CMD ["./myapp"]
