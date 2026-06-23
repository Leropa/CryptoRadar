package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type TickerRate struct {
	Exchange string
	Coin     string
	Price    float64
}

func ConnectDB() {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=leropa dbname=crypto_radar sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при открытии соединения: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Нет связи с базой данных: %v", err)
	}

	query := `CREATE TABLE IF NOT EXISTS crypto (
	id SERIAL PRIMARY KEY,
	exchange VARCHAR(50),
	coin VARCHAR(10), 
	price NUMERIC(18, 4),
	updated_at TIMESTAMP DEFAULT NOW()
	);`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatalf("Не удалось иницилизировать таблицу crypto: %v", err)
	}

	fmt.Println("БД успешно подключена и настроена!")
}

func SaveRate(exchange, coin string, price float64) {
	query := `INSERT INTO crypto (exchange, coin, price) VALUES ($1, $2, $3)`

	_, err := DB.Exec(query, exchange, coin, price)
	if err != nil {
		log.Printf("Не удалось сохранить курс в БД: %v\n", err)
	}
}

func GetLatestRates() ([]TickerRate, error) {

	query := `
		SELECT DISTINCT ON (exchange, coin) exchange, coin, price 
		FROM crypto 
		ORDER BY exchange, coin, updated_at DESC;
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rates []TickerRate

	for rows.Next() {
		var r TickerRate
		err := rows.Scan(&r.Exchange, &r.Coin, &r.Price)
		if err != nil {
			return nil, err
		}
		rates = append(rates, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rates, nil
}
