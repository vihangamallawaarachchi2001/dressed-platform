FROM golang:1.25

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

# Explicitly target linux/amd64 (most common for servers)
RUN go build -o gateway ./cmd/main.go && \
    chmod +x gateway

EXPOSE 8000

CMD ["./gateway"]