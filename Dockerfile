# Stage 1: Build the Go application
# Use a Go version that matches or exceeds the one in go.mod (go 1.24.4)
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/web

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/web/static ./web/static/

COPY --from=builder /app/web/template ./web/template/

COPY --from=builder /app/app .

CMD ["./app"]