FROM golang:latest as builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mytolet-api ./cmd/server

# Use a minimal alpine image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/mytolet-api .
COPY --from=builder /app/.env .

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./mytolet-api"]