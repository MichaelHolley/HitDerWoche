# Use a Go image as the base
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o backend .

# Set the entry point command to run the Go application
CMD ["./backend"]