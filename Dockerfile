# Use official Golang base image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy Go files
COPY . .

# Download dependencies
RUN go mod tidy

# Build the app
RUN go build -o jobqueue .

# Expose port
EXPOSE 8080

# Start app
CMD ["./jobqueue"]
