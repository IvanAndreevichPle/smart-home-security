# deployments/docker/Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/video-processor

EXPOSE 8080

CMD ["./main"]
