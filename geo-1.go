package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Location структура для хранения данных о геолокации
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func handleLocation(w http.ResponseWriter, r *http.Request) {
	// Парсинг данных о геолокации из тела запроса
	var loc Location
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loc)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
 }

 // Вывод данных о геолокации в консоль
 	fmt.Printf("Received location - Latitude: %f, Longitude: %f\n", loc.Latitude, loc.Longitude)
}

func main1() {
	// Установка обработчика запросов для пути "/location"
	http.HandleFunc("/location", handleLocation)

	// Запуск сервера на порту 8080
	fmt.Println("Server is on, listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}