FROM golang:1.24

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

# Explicitly target linux/amd64 (most common for servers)
RUN go build -o auth-service ./cmd/main.go && \
    chmod +x auth-service

EXPOSE 8001

CMD ["./auth-service"]