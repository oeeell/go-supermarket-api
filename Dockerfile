# Use Go 1.24 image
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o supermarket-api

# Use a smaller base image for final deployment
FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=builder /app/supermarket-api .

CMD ["/supermarket-api"]
