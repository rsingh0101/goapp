# Dockerfile for the Go-based mariadb crud app
FROM golang:1.22-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
COPY . .
COPY config.yaml /app/config.yaml
RUN go mod vendor
# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o mdb .

# Final stage
FROM golang:1.22-alpine

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/mdb .
COPY --from=builder /app/config.yaml /app/config.yaml
# # Command to run the binary
# CMD ["./consumer"]
