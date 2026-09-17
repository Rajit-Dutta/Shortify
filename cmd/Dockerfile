FROM golang:alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN go build -o main ./cmd/api

FROM alpine
RUN adduser -S -D -H -h appuser
USER appuser
WORKDIR /app
COPY --from=builder /build/main /app/main
EXPOSE 3000
CMD ["./main/"]