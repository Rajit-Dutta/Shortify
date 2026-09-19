FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
RUN go build -o main ./cmd/api

FROM alpine:latest
RUN adduser -S -D -H appuser
USER appuser
WORKDIR /app
COPY --from=builder /build/main /app/main
COPY cmd/.env /app/.env
EXPOSE 3000
CMD ["./main"]