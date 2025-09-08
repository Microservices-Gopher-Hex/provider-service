# ---- build stage ----
FROM golang:1.22 AS builder
WORKDIR /app

# mod cache
COPY go.mod go.sum ./
RUN go mod download

# código
COPY . .

# compilar binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/provider ./cmd/provider

# ---- final stage ----
FROM alpine:3.20
WORKDIR /app

# deps mínimas para esperar DB
RUN apk add --no-cache ca-certificates bash postgresql-client curl

# binario y scripts
COPY --from=builder /out/provider /app/provider
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# variables de entorno
ENV HTTP_PORT=8082 \
    DB_DSN="" \
    KAFKA_BROKERS=""

EXPOSE 8082
ENTRYPOINT ["/app/entrypoint.sh"]