package main

import (
	"cryptoradar/handlers"
	"fmt"
	"net/http"
	"time"
)

func main() {

	go func() {
		for {
			fmt.Println("[Сборщик]: Свежие курсы успешно запрошены...")

			time.Sleep(5 * time.Second)
		}
	}()

	router := handlers.NewRouter()

	fmt.Println("🚀 Сервер успешно запущен!")
	fmt.Println("🌍 Открывай в браузере: http://localhost:8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("❌ Ошибка запуска сервера:", err)
	}
}
