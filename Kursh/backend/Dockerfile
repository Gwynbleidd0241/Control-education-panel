# ---- build stage ----
FROM golang:1.25 AS builder

WORKDIR /app

# сначала зависимости
COPY go.mod go.sum ./
RUN go mod download

# затем код
COPY . .

# собираем бинарь
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

# ---- run stage ----
FROM alpine:3.20

RUN adduser -D -g '' appuser

WORKDIR /app

# копируем бинарь и миграции (если они нужны в контейнере)
COPY --from=builder /app/app /app/app

ENV PORT=8081

USER appuser

EXPOSE 8081

CMD ["/app/app"]
