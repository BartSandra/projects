FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

# Установка инструмента для генерации документации Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/api/main.go

RUN go build -o main cmd/api/main.go
RUN apk --no-cache add postgresql-client

COPY wait-for-postgres.sh /wait-for-postgres.sh
RUN chmod +x /wait-for-postgres.sh

EXPOSE 8080

CMD ["/app/main"]
