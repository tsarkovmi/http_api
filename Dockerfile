FROM golang:1.22.5

# Установка значения для переменной окружения
RUN go version 
ENV GOPATH=/

# Копируем все файлы из корневой директории проекта
COPY ./ ./

# Устанавливаем psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# Делаем скрип для ожидания в виде запускаемой программы
#RUN chmod +x wait-for-postgres.sh

# Скачивание зависимостей и компиляция 
# приложения через go build
RUN go mod download
RUN go build -o httpapi ./cmd/main.go
CMD ["./httpapi"]