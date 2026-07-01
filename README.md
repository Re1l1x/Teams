# Teams REST API

Тестовое задание на Golang:
REST API сервис для управления группами людей.

Выполнил:
Дмитрий Дьячков

## Используемые технологии

- Go 1.26.1
- Gin
- PostgreSQL 17
- Docker & Docker Compose

## Запуск проекта

### 1. Запустить контейнеры

```bash
docker compose up --build -d
```

### 2. Применить миграции

```bash
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5433/people_groups?sslmode=disable" up
```

### 3. API будет доступен по адресу

```text
http://localhost:8080
```

## Основные маршруты

### Группы

| Метод  | URL                       | Описание                                     |
| ------ | ------------------------- | -------------------------------------------- |
| POST   | `/groups`                 | Создать группу                               |
| GET    | `/groups`                 | Получить список групп                        |
| GET    | `/groups/{id}`            | Получить группу по ID                        |
| PUT    | `/groups/{id}`            | Обновить группу                              |
| DELETE | `/groups/{id}`            | Удалить группу                               |
| GET    | `/groups/{id}/people`     | Люди только данной группы                    |
| GET    | `/groups/{id}/people/all` | Люди группы и всех дочерних групп            |
| GET    | `/groups/{id}/count`      | Количество людей только в группе             |
| GET    | `/groups/{id}/count/all`  | Количество людей в группе и дочерних группах |

### Люди

| Метод  | URL            | Описание                |
| ------ | -------------- | ----------------------- |
| POST   | `/people`      | Создать человека        |
| GET    | `/people`      | Получить список людей   |
| GET    | `/people/{id}` | Получить человека по ID |
| PUT    | `/people/{id}` | Обновить человека       |
| DELETE | `/people/{id}` | Удалить человека        |

## Миграции

Применить:

```bash
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5433/people_groups?sslmode=disable" up
```

Откатить:

```bash
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5433/people_groups?sslmode=disable" down -all
```
