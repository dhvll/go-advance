# Build 
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .

# Set entrypoint and command
EXPOSE 8080 
CMD [ "/app/main" ]