# Build stage
FROM golang:1.19-alpine as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o app

# Production stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
