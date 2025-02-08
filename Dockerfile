# Use official Golang image as base
FROM golang:1.23

# Set working directory
WORKDIR /app

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the application
RUN go build -o receipt-processor main.go

# Expose port 8080 for API requests
EXPOSE 8080

# Start the server
CMD ["./receipt-processor"]
