# Erste Build-Stage: Kompiliere den Code für AMD64
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o main

# Zweite Build-Stage: Erstelle das AMD64/Linux-Image
FROM amd64/alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
RUN chmod +x /app/main
EXPOSE 5080
CMD ["./main"]