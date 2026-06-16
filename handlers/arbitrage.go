package handlers

import (
	"net/http"
)

// LiveRateResponse описывает формат данных для одной монеты.
// Теги json:"..." обязательны, чтобы JavaScript в app.js смог прочитать поля.
type LiveRateResponse struct {
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	BinancePrice float64 `json:"binance_price"`
	BybitPrice   float64 `json:"bybit_price"`
	BestBuyAt    string  `json:"best_buy_at"`
	Spread       float64 `json:"spread"`
}

// GetArbitrageData формирует список монет и отправляет его на фронтенд
func GetArbitrageData(w http.ResponseWriter, r *http.Request) {
	// Создаем слайс (массив) для хранения результатов
	var results []LiveRateResponse

	// 1. Забиваем данные для Bitcoin
	btc := LiveRateResponse{
		Name:         "Bitcoin",
		Symbol:       "BTC",
		BinancePrice: 65000.0,
		BybitPrice:   65200.0,
		BestBuyAt:    "Binance",
		Spread:       0.31,
	}
	results = append(results, btc)

	// 2. Забиваем данные для Ethereum
	eth := LiveRateResponse{
		Name:         "Ethereum",
		Symbol:       "ETH",
		BinancePrice: 3515.0,
		BybitPrice:   3500.0,
		BestBuyAt:    "Bybit",
		Spread:       0.43,
	}
	results = append(results, eth)

	// 3. Забиваем данные для Solana
	sol := LiveRateResponse{
		Name:         "Solana",
		Symbol:       "SOL",
		BinancePrice: 145.5,
		BybitPrice:   147.0,
		BestBuyAt:    "Binance",
		Spread:       1.03,
	}
	results = append(results, sol)

	// Отправляем готовый слайс через дженерик-функцию
	RespondWithJSON(w, http.StatusOK, results)
}
