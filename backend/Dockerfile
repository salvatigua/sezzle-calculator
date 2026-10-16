# Stage 1: Build (usamos alpine para incluir la versión más reciente de Go)
FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Stage 2: Minimal Runtime Image
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8081
CMD ["./main"]