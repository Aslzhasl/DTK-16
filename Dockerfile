# Стадия сборки
FROM golang:1.20 AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка бинарника
RUN go build -o main ./cmd/main.go

# Стадия запуска
FROM debian:bullseye-slim

# Создаём директорию приложения
WORKDIR /app

# Копируем бинарник из стадии builder
COPY --from=builder /app/main .

# Копируем .env файл
COPY .env .

# Открываем порт
EXPOSE 8080

# Запуск приложения
CMD ["./main"]
