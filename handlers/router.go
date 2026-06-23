package handlers

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON[T any](w http.ResponseWriter, statusCode int, data T) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)

	mux.HandleFunc("/api/arbitrage", GetArbitrageData)

	return mux
}
