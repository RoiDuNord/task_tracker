# Task Tracker

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Chi](https://img.shields.io/badge/chi-%23000000.svg?style=for-the-badge&logo=&logoColor=white)

---

## Описание проекта

Task Tracker на Golang с горутинами, многопоточностью с использованием роутера Chi, реализующий RESTful API.

---

## Особенности и функциональные требования

- **Роутинг и API:**
```
POST  http://localhost:8080/tasks
GET  http://localhost:8080/tasks/{id}
DELETE  http://localhost:8080/tasks/{id}

```

## Установка и запуск

1. Клонируйте репозиторий

```
git clone https://github.com/RoiDuNord/linked_list_cache.git
```

2. Запустите программу

```
go run cmd/main.go
```

3. Передайте задачу с помощью json

```
POST  http://localhost:8080/tasks

{
    "taskName": "sleep"
}
```

4. Проверьте состояние задачи

```
GET  http://localhost:8080/tasks/{id}
```

5. Удалите задачу

```
DELETE  http://localhost:8080/tasks/{id}
```