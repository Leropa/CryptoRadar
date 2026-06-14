package fetchers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type RateFetcher interface {
	FetchPrice(coin string) (float64, error)
}

type BinanceFetcher struct{}

type BinanceResponse struct {
	Price string `json:"price"`
}

type BybitFetcher struct{}

type BybitResponse struct {
	Result struct {
		List []struct {
			LastPrice string `json:"lastPrice"`
		} `json:"list"`
	} `json:"result"`
}

func (bi BinanceFetcher) FetchPrice(coin string) (float64, error) {
	request := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", coin)

	resl, err := http.Get(request)
	if err != nil {
		fmt.Println("Не получилось получить данные ", err)
	}
	defer resl.Body.Close()

	var biRes BinanceResponse

	err = json.NewDecoder(resl.Body).Decode(&biRes)
	if err != nil {
		return 0, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	priceFloat, err := strconv.ParseFloat(biRes.Price, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка конвертации цены в число: %v", err)
	}

	return priceFloat, nil

}

func (by BybitFetcher) FetchPrice(coin string) (float64, error) {
	request := fmt.Sprintf("https://api.bybit.com/v5/market/tickers?category=linear&symbol=%sUSDT", coin)

	resl, err := http.Get(request)
	if err != nil {
		return 0, fmt.Errorf("ошибка HTTP-запроса: %v", err)
	}
	defer resl.Body.Close()

	var byRes BybitResponse
	err = json.NewDecoder(resl.Body).Decode(&byRes)
	if err != nil {
		return 0, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	priceFloat, err := strconv.ParseFloat(byRes.Result.List[0].LastPrice, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка конвертации цены в число: %v", err)
	}

	return priceFloat, nil
}
