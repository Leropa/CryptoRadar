package handlers

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON — дженерик-функция для быстрой отправки JSON на фронтенд.
// Мы используем [T any], чтобы передавать сюда любые структуры данных.
func RespondWithJSON[T any](w http.ResponseWriter, statusCode int, data T) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Настраиваем раздачу фронтенда (папка public) по корневому пути "/"
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)

	// Регистрируем наш API-эндпоинт для выдачи курсов крипты
	mux.HandleFunc("/api/arbitrage", GetArbitrageData)

	return mux
}
