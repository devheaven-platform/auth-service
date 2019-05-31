# Builder
FROM golang:1.12.5 as builder

# Set working directory
WORKDIR /app

# Copy mod & sum files
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

# Install dependencies
RUN go mod download

# Copy source
COPY . /app

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/bin/auth-service /app/cmd/api/main.go

# Worker
FROM scratch

# Set working directory
WORKDIR /app/bin

# Copy binary & swagger
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/auth-service /app/bin/auth-service
COPY --from=builder /app/spec /app/bin/spec
COPY --from=builder /app/dist /app/bin/dist

# Run binary
ENTRYPOINT ["./auth-service"]