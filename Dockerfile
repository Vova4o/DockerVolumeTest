FROM alpine:3.12

WORKDIR /app

# Установка bash и других необходимых утилит
RUN apk add --no-cache bash

# Копируем скомпилированное приложение из локальной системы
COPY maintest /app

# Устанавливаем права на выполнение для приложения
RUN chmod +x /app/maintest

CMD ["./maintest"]
