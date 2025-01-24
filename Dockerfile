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

# Update and install PostgreSQL client and netcat-openbsd dependencies
RUN apt-get update && apt-get install -y postgresql-client netcat-openbsd

# Install Redis client dependencies
RUN go get github.com/redis/go-redis/v9

# Copy the wait-for-it.sh script to the working directory
COPY scripts/wait-for-it.sh /app/wait-for-it.sh

# Make the script executable
RUN chmod +x /app/wait-for-it.sh

# Build the Go application
RUN go build

# Expose the port your application listens on
EXPOSE 8080

# Command to run the application
CMD [ "./wait-for-it.sh", "postgres:5432", "--", "./wait-for-it.sh", "redis:6379", "--", "./reusable-api"]
