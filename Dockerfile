# Stage 1: Build the binary
FROM golang:1.27-alpine AS builder
WORKDIR /app

# Copy go module files first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

RUN go generate .
# Build a statically linked binary
# CGO_ENABLED=0 ensures no C dependencies
# -ldflags="-s -w" strips debug info to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o /app/server \
    ./cmd/server

##

# Stage 2: Create the minimal runtime image
FROM scratch

# Copy the binary
COPY --from=builder /app/server /server

# Cloud Run listens on this port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/server"]
