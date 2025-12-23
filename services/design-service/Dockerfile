FROM golang:1.25

WORKDIR /app

RUN mkdir -p uploads

COPY go.* ./
RUN go mod download

COPY . .

# Explicitly target linux/amd64 (most common for servers)
RUN go build -o design-service ./cmd/main.go && \
    chmod +x design-service

EXPOSE 8002

CMD ["./design-service"]