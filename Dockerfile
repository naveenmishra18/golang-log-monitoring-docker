
# Stage 1: Build the Go application
FROM golang:1.20 AS build_stage
WORKDIR /src

# Copy module files and resolve dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy source code and build the binary
COPY . .
RUN go build -o /src/server main.go  # Build the executable

# Stage 2: Create minimal runtime image
FROM alpine:3.18
WORKDIR /app

# Install necessary utilities
RUN apk add --no-cache logrotate

# Copy the built application from the previous stage
COPY --from=build_stage /src/server .

# Set executable permissions
RUN chmod +x /app/server

# Create log directory
RUN mkdir -p /var/log/app_logs

# Define the container startup command
CMD ["/app/server"]
