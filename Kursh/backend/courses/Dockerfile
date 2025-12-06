FROM golang:1.25-alpine AS builder
WORKDIR /src

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" \
    -o /app/courses ./cmd/courses

FROM alpine:3.20
WORKDIR /app
RUN adduser -D app
USER app

COPY --from=builder /app/courses /app/courses
EXPOSE 8080
ENTRYPOINT ["/app/courses"]
