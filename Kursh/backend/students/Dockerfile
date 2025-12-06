FROM golang:1.25-alpine AS builder
WORKDIR /src

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" \
    -o /app/students ./cmd/students

FROM alpine:3.20
WORKDIR /app
RUN adduser -D app
USER app

COPY --from=builder /app/students /app/students
EXPOSE 8080
ENTRYPOINT ["/app/students"]
