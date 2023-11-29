package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GeoData1 struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func ProcessGeoData(data GeoData) string {
	result := fmt.Sprintf("Processed GeoData: Latitude %f, Longitude %f", data.Latitude, data.Longitude)
	return result
}

func HandleGeoData(w http.ResponseWriter, r *http.Request) {
	var geoData GeoData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&geoData)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	result := ProcessGeoData(geoData)
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func main2() {
	http.HandleFunc("/process-geo", HandleGeoData)
	fmt.Println("Server is on, listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
