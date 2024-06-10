# Base stage
FROM golang:1.22.3-alpine AS base

# Add Maintainer Info
LABEL maintainer="Nguyen Tien Tai <nguyentientai10@gmail.com>"

# Install dependencies
RUN apk add --no-cache curl

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# *-------------- Build Server --------------*
FROM base AS server-builder

# Build the Go server app
RUN go build -o /server ./cmd/server/main.go

# *-------------- Build Cron Job --------------*
FROM base AS cron-builder

# Build the Go cron job app
RUN go build -o /cron ./cmd/cronjob/main.go

# *-------------- Final Stage for Server --------------*
FROM debian:stable-slim AS server

# Copy the server binary from the builder stage
COPY --from=server-builder /server /app/server

# Expose port 8080 to the outside world (for server)
EXPOSE 8080

# Command to run the server executable
CMD ["/app/server"]

# *-------------- Final Stage for Cron Job --------------*
FROM debian:stable-slim AS cron

# Copy the cron job binary from the builder stage
COPY --from=cron-builder /cron /app/cron

# Command to run the cron executable
CMD ["/app/cron"]
