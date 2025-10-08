# User Service API

Базовый путь: http://{HOST}:{PORT}/api

По умолчанию в проекте сервис слушает порт 8081 (переменная окружения `PORT`).

## Аутентификация

При успешном логине сервис возвращает JWT токен. Для защищённых эндпоинтов нужно передавать заголовок:

Authorization: Bearer <token>

Токен подписывается секретом из `JWT_SECRET` (переменная окружения).

## Эндпоинты

1) Регистрация пользователя

- URL: POST /api/users/register
- Описание: создает нового пользователя.
- Тело запроса (JSON):

```json
{
  "email": "user@example.com",
  "password": "strong_password",
  "name": "User Name"
}
```

- Успешный ответ: 201 Created

Пример тела ответа (объект User):

```json
{
  "id": "uuid",
  "email": "user@example.com",
  "name": "User Name"
}
```

- Ошибки:
  - 400 Bad Request — неверный формат запроса
  - 409 Conflict — пользователь с таким email уже существует

2) Логин

- URL: POST /api/users/login
- Описание: проверяет email/password и возвращает JWT токен и информацию о пользователе.
- Тело запроса (JSON):

```json
{
  "email": "user@example.com",
  "password": "strong_password"
}
```

- Успешный ответ: 200 OK

Пример тела ответа:

```json
{
  "token": "<jwt-token>",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "User Name"
  }
}
```

- Ошибки:
  - 400 Bad Request — неверный формат запроса
  - 401 Unauthorized — неправильные учетные данные


- Ошибки:
  - 404 Not Found — пользователь с таким id не найден


## Примеры curl

Регистрация:

```bash
curl -s -X POST http://localhost:8081/api/users/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","name":"Test"}'
```

Логин:

```bash
curl -s -X POST http://localhost:8081/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

Получить пользователя:

```bash
curl -s http://localhost:8081/api/users/<user-id>
```

Пример с авторизацией (если требуется токен):

```bash
curl -s http://localhost:8081/api/users/<user-id> \
  -H "Authorization: Bearer <jwt-token>"
```


## Переменные окружения (важные)

- PORT — порт, на котором запускается сервис (по умолчанию 8081)
- DB_HOST — хост PostgreSQL (в compose: `postgres-users`)
- DB_PORT — порт PostgreSQL (по умолчанию 5432)
- DB_USER — пользователь БД (пример: `trello_user`)
- DB_PASSWORD — пароль БД
- DB_NAME — имя базы данных (пример: `users`)
- JWT_SECRET — секрет для подписи JWT


## Запуск локально (через docker-compose)

В каталоге `server`:

```bash
docker-compose up -d --build
```

После запуска сервис будет доступен по адресу http://localhost:8081
