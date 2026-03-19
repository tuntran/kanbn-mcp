FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /kanbn-mcp ./cmd/server

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
COPY --from=builder /kanbn-mcp /kanbn-mcp
# MCP_HTTP_ADDR defaults to :8080 for container deployments (Coolify, Docker).
# Unset or override to use stdio transport instead.
ENV MCP_HTTP_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/kanbn-mcp"]
