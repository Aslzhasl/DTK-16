FROM golang:1.23.4

WORKDIR /app

# Загрузка зависимостей
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копирование проекта
COPY . .

# Скрипт ожидания базы
#COPY wait-for-postgres.sh /wait-for-postgres.sh
#RUN chmod +x /wait-for-postgres.sh

# Сборка
RUN go build -o app ./cmd

EXPOSE 8082

CMD [ "./app"]
