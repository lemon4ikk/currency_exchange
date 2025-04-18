# Currency Exchange

Проект реализует простой REST API для конвертации валют на основе заранее заданных курсов обмена.

## Функционал

- Получение списка всех доступных валют
- Получение курса обмена между двумя валютами
- Конвертация суммы из одной валюты в другую

## Технологии

- Go 1.24
- Go Modules
- Docker

## Структура проекта

- `cmd/main.go`: точка входа
- `api/`: 
- `handler/`: обработчики HTTP-запросов
- `repository/`: хранение валют и курсов (In-Memory)
- `service/`: бизнес-логика работы с валютами и курсами
- `templates`: 

## API

#### Получить список всех валют

`GET /currencies`

---

#### Получить курс обмена между двумя валютами

`GET /exchange-rate?from={fromCurrency}&to={toCurrency}`

**Параметры:**
- `from`: код исходной валюты
- `to`: код целевой валюты

## Деплой

`http://ip:8080/`

## Запуск проекта локально

1. Клонировать репозиторий:
    ```bash
    git clone https://github.com/your-username/currency-exchange.git
    cd currency-exchange
    ```

2. Установить зависимости:
    ```bash
    go mod tidy
    ```

3. Запустить сервер:
    ```bash
    go run main.go
    ```

4. API будет доступен на `http://localhost:8080`.