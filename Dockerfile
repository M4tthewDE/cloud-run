FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build cmd/server/main.go

From alpine:latest
LABEL org.opencontainers.image.source="https://github.com/m4tthewde/cloud-run"

RUN apk add -U libcap

WORKDIR /root/
COPY --from=builder /app/main ./
CMD ["./main"]
