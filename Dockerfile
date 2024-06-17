# Build stage
FROM golang:1.22 AS builder

# Download dependencies
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and compile it
COPY . .
RUN go build -o /app/ ./...

# Production stage - final image
FROM debian:bookworm-slim

WORKDIR /app
EXPOSE 8080/tcp 8081/tcp

# Install necessary runtime libraries
RUN apt-get update \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app /app

CMD ["/app/server"]