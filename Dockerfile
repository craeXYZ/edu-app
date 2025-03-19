# Use an official Golang image as the base
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN go build -o main .

# Expose the API port
EXPOSE 8080

# Run the application
CMD ["/app/main"]
