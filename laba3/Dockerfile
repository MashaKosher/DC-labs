# Указываем базовый образ
FROM golang:1.20

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Компилируем приложение
RUN go build -o main .

# Запускаем приложение
CMD ["./main"]