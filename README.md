# Currency Exchange

Проект реализует простой REST API для конвертации валют на основе заранее заданных курсов обмена.

## Установка и запуск проекта локально

1. Убедитесь, что у вас установлен Go.

2. Склонируйте репозиторий и перейдите в рабочую директорию:

   ```sh
   git clone https://github.com/lemon4ikk/currency_exchange.git
   cd currency_exchange
   ```

3. Запустите проект:

   ```sh
   go run cmd/main.go
   ```

4. API будет доступен на `http://localhost:8080`

## API

## Валюты

### Получение списка всех валют
#### GET `/currencies`

### Получение конкретной валюты
#### GET `/currency/EUR`

### Добавление новой валюты в базу
#### POST `/currencies`
#### Данные передаются в теле запроса в виде полей формы (x-www-form-urlencoded). Поля формы - name, code, sign.
---
## Обменные курсы

### Получение списка всех валют
#### GET `/exchangeRates`

### Получение конкретного обменного курса
#### GET `GET /exchangeRate/USDRUB`

### Добавление новой валюты в базу
#### POST `/exchangeRates`
#### Данные передаются в теле запроса в виде полей формы (x-www-form-urlencoded). Поля формы - baseCurrencyCode, targetCurrencyCode, rate.

### Обновление существующего в базе обменного курса
#### PATCH `/exchangeRate/USDRUB`
#### Данные передаются в теле запроса в виде полей формы (x-www-form-urlencoded). Единственное поле формы - rate
---

## Обмен валюты

### Расчёт перевода определённого количества средств из одной валюты в другую
#### GET `/exchange?from=BASE_CURRENCY_CODE&to=TARGET_CURRENCY_CODE&amount=$AMOUNT`
#### Пример запроса - GET /exchange?from=USD&to=AUD&amount=10