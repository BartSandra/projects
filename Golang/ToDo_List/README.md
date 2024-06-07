# Todo List Application

## Описание

Это приложение представляет собой API для управления списком задач (Todo List). Приложение написано на Go и использует PostgreSQL в качестве базы данных.

## Требования

Для запуска проекта необходимо установить:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Установка и запуск

1. Клонируйте репозиторий:

Создайте и запустите контейнеры Docker:
```
docker-compose up --build
```
Это создаст и запустит контейнеры для приложения и базы данных.

Приложение будет доступно по адресу http://localhost:8080.

## Использование
Приложение предоставляет следующие эндпоинты:

GET /todos - Получить список всех задач
GET /todos/{id} - Получить задачу по ID
POST /todos - Создать новую задачу
PUT /todos/{id} - Обновить задачу по ID
DELETE /todos/{id} - Удалить задачу по ID

## Примеры запросов
Получить список всех задач:
```
curl -X GET http://localhost:8080/todos
```
Создать новую задачу:
```
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title": "New Task", "description": "Task description"}'
```
## Тестирование
Для запуска тестов используйте следующую команду:
```
docker-compose run --rm test
```
## API Документация

Swagger-документация доступна по адресу `http://localhost:8080/swagger/`.
