FROM golang:1.25

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

# Explicitly target linux/amd64 (most common for servers)
RUN go build -o supplier-service ./cmd/main.go && \
    chmod +x supplier-service

EXPOSE 8002

CMD ["./supplier-service"]