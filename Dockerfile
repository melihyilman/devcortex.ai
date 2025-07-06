# Stage 1: Build the Go application
# Use a Go version that matches or exceeds the one in go.mod (go 1.24.4)
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application, creating a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/web

# Stage 2: Create the final, minimal image
FROM alpine:latest

COPY --from=builder /app/app .

# The command to run when the container starts.
CMD ["./app"]