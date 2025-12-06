FROM golang:1.25-alpine AS builder
WORKDIR /src

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" \
    -o /app/gateway ./cmd/gateway

FROM alpine:3.20
WORKDIR /app
RUN adduser -D app
USER app

COPY --from=builder /app/gateway /app/gateway
EXPOSE 8080
ENTRYPOINT ["/app/gateway"]
