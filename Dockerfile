FROM golang:alpine AS builder
ENV GOTOOLCHAIN=auto
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.18
ENV GOTOOLCHAIN=auto
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8087
CMD ["./main"]
