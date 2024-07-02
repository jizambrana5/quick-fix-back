# Use an official Golang runtime as a parent image
FROM golang:1.21.3

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY go.mod go.sum ./

# Install Go dependencies
RUN go mod download

# Copy the rest of your application code
COPY . .

COPY .env /app/pkg

# Navigate to the directory containing the main file
WORKDIR /app/pkg

# Build the Go application
RUN go build -o /app/myapp

# Expose a port for the application to listen on
EXPOSE 8080

# Run the application
CMD ["/app/myapp"]
