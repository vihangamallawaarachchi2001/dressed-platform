FROM golang:1.25

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

# Explicitly target linux/amd64 (most common for servers)
RUN go build -o payment-service ./cmd/main.go && \
    chmod +x payment-service

EXPOSE 8002

CMD ["./payment-service"]