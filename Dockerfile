# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/cost-optimizer ./cmd

# Runtime stage
FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app

# Run as non-root for better container security.
RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /app/cost-optimizer ./cost-optimizer
COPY --from=builder /src/web ./web
RUN chown -R app:app /app

USER app

EXPOSE 8080
ENTRYPOINT ["./cost-optimizer"]
