# CryptoRadar

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/Tailwind_CSS-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

**Мониторинг криптовалютных курсов и арбитражных возможностей в реальном времени**

CryptoRadar сравнивает цены на Binance и Bybit, показывает лучшую биржу для покупки и размер спреда — всё в одном дашборде.

---

## Скриншот

```
┌─────────────────────────────────────────────────────────────┐
│  CryptoRadar v1.0          ◉ Live API                      │
│  Мониторинг лучшего курса и арбитраж в реальном времени    │
├──────────────────┬──────────────────┬───────────────────────┤
│  Bitcoin (BTC)   │  Ethereum (ETH)  │  Solana (SOL)         │
│                  │                  │                       │
│  Binance: $65000 │  Binance: $3515  │  Binance: $145.50     │
│  Bybit:   $65200 │  Bybit:   $3500  │  Bybit:   $147.00     │
│                  │                  │                       │
│  Лучшая сделка   │  Лучшая сделка   │  Лучшая сделка        │
│  Купить: Binance │  Купить: Bybit   │  Купить: Binance      │
│  Спред: +0.31%   │  Спред: +0.43%   │  Спред: +1.03%        │
└──────────────────┴──────────────────┴───────────────────────┘
```

---

## Возможности

- Сравнение цен между **Binance** и **Bybit** в реальном времени
- Автоматическое определение **лучшей биржи** для покупки
- Расчёт **спреда** между биржами (%)
- Красивый **dashboard** с glassmorphism-эффектом
- Автообновление данных каждую секунду
- Хранение истории курсов в **PostgreSQL**

---

## Структура проекта

```
CryptoRadar/
├── cmd/
│   └── main.go          # Точка входа, запуск HTTP-сервера
├── db/
│   └── db.go            # Подключение к PostgreSQL, работа с БД
├── fetchers/
│   └── fetcher.go       # Получение цен с бирж (Binance, Bybit)
├── handlers/
│   ├── router.go        # Маршрутизация и JSON-ответы
│   └── arbitrage.go     # Логика арбитража и API-эндпоинт
├── public/
│   ├── index.html       # Frontend-дашборд
│   ├── css/style.css    # Кастомные стили (glassmorphism)
│   └── js/app.js        # Запросы к API и отрисовка карточек
├── go.mod               # Модуль Go и зависимости
└── .gitignore
```

---

## Технологии

| Слой | Технологии |
|------|-----------|
| **Backend** | Go 1.26, net/http, encoding/json |
| **Database** | PostgreSQL 15+, lib/pq |
| **Frontend** | HTML5, Tailwind CSS, Vanilla JavaScript |
| **API** | Binance API, Bybit API |

---

## Быстрый старт

### 1. Клонирование репозитория

```bash
git clone https://github.com/Leropa/CryptoRadar.git
cd CryptoRadar
```

### 2. Установка зависимостей

```bash
go mod download
```

### 3. Настройка базы данных

Убедись, что PostgreSQL запущен. Создай базу данных:

```sql
CREATE DATABASE crypto_radar;
```

Отредактируй строку подключения в `db/db.go`:

```go
connStr := "host=localhost port=5432 user=postgres password=YOUR_PASSWORD dbname=crypto_radar sslmode=disable"
```

### 4. Запуск

```bash
go run cmd/main.go
```

Сервер запустится на `http://localhost:8080`. Открой эту ссылку в браузере.

---

## API

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| `GET` | `/` | Статический frontend-дашборд |
| `GET` | `/api/arbitrage` | JSON с арбитражными данными по всем монетам |

### Пример ответа `/api/arbitrage`

```json
[
  {
    "name": "Bitcoin",
    "symbol": "BTC",
    "binance_price": 65000.0,
    "bybit_price": 65200.0,
    "best_buy_at": "Binance",
    "spread": 0.31
  }
]
```

---

## Roadmap

- [ ] Подключить реальные данные вместо моков в обработчике
- [ ] Добавить больше бирж (OKX, Huobi, Gate.io)
- [ ] Реализовать историю цен с графиками
- [ ] Уведомления при большом спреде (Telegram-бот)
- [ ] Добавить аутентификацию пользователей
- [ ] Docker-контейнер для быстрого развертывания

---

## Лицензия

MIT License - свободное использование и модификация.

---

*Создано для отслеживания арбитражных возможностей на криптовалютном рынке*
