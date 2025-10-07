# Документация API

# Архитектура
Проект разделен на 3 микросервиса:
![arch](/docs/arch/arch.png)

## User Service
Сервис управления пользователями и аутентификацией

### Эндпоинты:
- POST /users/register - Регистрация нового пользователя
- POST /users/login - Авторизация пользователя

## Board Service
Сервис управления досками и колонками

### Эндпоинты:
- POST /boards - Создание новой доски
- GET /boards - Получение списка досок пользователя
- GET /boards/{id} - Получение информации о доске
- PUT /boards/{id} - Обновление доски
- DELETE /boards/{id} - Удаление доски
- POST /boards/{id}/columns - Создание колонки в доске
- GET /boards/{id}/columns - Получение колонок доски
- PUT /boards/{id}/columns/{columnId} - Обновление колонки
- DELETE /boards/{id}/columns/{columnId} - Удаление колонки
- POST /boards/{id}/members - Добавление участника в доску
- DELETE /boards/{id}/members/{userId} - Удаление участника из доски

## Task Service
Сервис управления задачами (карточками)

### Эндпоинты:
- POST /tasks - Создание новой задачи
- GET /tasks/board/{boardId} - Получение задач доски
- GET /tasks/{id} - Получение информации о задаче
- PUT /tasks/{id} - Обновление задачи
- DELETE /tasks/{id} - Удаление задачи
- PATCH /tasks/{id}/move - Перемещение задачи между колонками
- PATCH /tasks/{id}/assign - Назначение исполнителя задачи
- POST /tasks/{id}/comments - Добавление комментария к задаче
- GET /tasks/{id}/comments - Получение комментариев задачи
