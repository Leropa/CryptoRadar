package handlers

import (
	"net/http"
)

type LiveRateResponse struct {
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	BinancePrice float64 `json:"binance_price"`
	BybitPrice   float64 `json:"bybit_price"`
	BestBuyAt    string  `json:"best_buy_at"`
	Spread       float64 `json:"spread"`
}

func GetArbitrageData(w http.ResponseWriter, r *http.Request) {
	var results []LiveRateResponse

	btc := LiveRateResponse{
		Name:         "Bitcoin",
		Symbol:       "BTC",
		BinancePrice: 65000.0,
		BybitPrice:   65200.0,
		BestBuyAt:    "Binance",
		Spread:       0.31,
	}
	results = append(results, btc)

	eth := LiveRateResponse{
		Name:         "Ethereum",
		Symbol:       "ETH",
		BinancePrice: 3515.0,
		BybitPrice:   3500.0,
		BestBuyAt:    "Bybit",
		Spread:       0.43,
	}
	results = append(results, eth)

	sol := LiveRateResponse{
		Name:         "Solana",
		Symbol:       "SOL",
		BinancePrice: 145.5,
		BybitPrice:   147.0,
		BestBuyAt:    "Binance",
		Spread:       1.03,
	}
	results = append(results, sol)

	RespondWithJSON(w, http.StatusOK, results)
}
