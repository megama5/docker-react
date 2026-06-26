FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/simple_web.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]