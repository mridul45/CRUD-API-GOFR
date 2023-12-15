# Use a larger Golang image for downloading dependencies
FROM golang:1.21.4 AS builder

WORKDIR /app

# Copy only the go.mod and go.sum files to leverage caching
COPY go.mod go.sum ./

# Download dependencies using go modules
RUN go mod download

# Copy the entire application (including Go source files)
COPY . .

# Build the Go application with CGO_ENABLED=0 to create a statically-linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Build stage: Use distroless for the final image
FROM gcr.io/distroless/base:latest

WORKDIR /app

# Copy the main binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose the port that the application will run on
EXPOSE 8000

# Command to run the executable
CMD ["/app/main"]
