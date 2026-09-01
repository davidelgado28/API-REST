FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o task-api main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/task-api .
EXPOSE 8080
CMD ["./task-api"]
