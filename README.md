# Калькулятор выражений с REST API

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yourusername/calculator)
![License](https://img.shields.io/badge/license-MIT-blue)

Простой микросервис для вычисления математических выражений с аутентификацией пользователей и историей операций.

## 📋 Содержание
- [Особенности](#✨-особенности)
- [Требования](#⚙️-требования)
- [Установка](#🚀-установка)
- [Использование API](#🔧-использование-api)
- [Примеры](#📚-примеры)
- [Разработка](#👨‍💻-разработка)
- [Лицензия](#📜-лицензия)

## ✨ Особенности
- ✅ Вычисление сложных математических выражений
- 🔐 JWT аутентификация
- 💾 Сохранение истории вычислений в SQLite
- 📊 Поддержка математических функций
- ⚡ Быстрая обработка запросов

## ⚙️ Требования
- Go 1.20+
- SQLite3
- Git (для установки)

## 🚀 Установка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/yourusername/calculator.git
cd calculator
Установите зависимости:

bash
go mod download
Запустите сервер:

bash
go run main.go
Сервер запустится на http://localhost:8080

🔧 Использование API
🔐 Аутентификация
Регистрация:

bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser", "password":"testpassword123"}'
Вход:

bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser", "password":"testpassword123"}'
🧮 Вычисления
Отправить выражение:

bash
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"expression":"2 + 2 * 2"}'
Получить историю:

bash
curl -X GET http://localhost:8080/tasks \
  -H "Authorization: Bearer YOUR_TOKEN"
📚 Примеры выражений
Выражение	Результат
3 + 4 * 2	11
(1 + 2) * 3	9
sqrt(16) + pow(2, 3)	12
sin(3.1415926535/2)	1
👨‍💻 Разработка
Структура проекта:
calculator/
├── api/                 # Protobuf определения
├── internal/            # Внутренние пакеты
│   ├── calculator/      # Логика вычислений
│   ├── grpcserver/      # gRPC сервер
│   └── storage/         # Работа с базой данных
├── main.go              # Точка входа
└── README.md            # Документация
Сборка:
bash
go build -o calculator
Тестирование:
bash
go test ./...
📜 Лицензия
Этот проект распространяется под лицензией MIT. См. файл LICENSE для получения дополнительной информации.