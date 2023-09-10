# Use the official Golang image as the base image
FROM golang

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Install PostgreSQL client dependencies
RUN apt-get update && apt-get install -y postgresql-client

# Install Redis client dependencies
RUN go get github.com/redis/go-redis/v9

# Build the Go application
RUN go build

# Expose the port your application listens on
EXPOSE 8080

# Command to run the application
CMD ["postgres:5432", "./reusable-api"]

