package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MungaSoftwiz/location-web-server/internal/services"
	"github.com/MungaSoftwiz/location-web-server/internal/utils"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	visitorName := r.URL.Query().Get("visitor_name")
	if visitorName == "" {
		visitorName = "Guest"
	}

	clientIP := utils.GetClientIP(r)
	location, err := services.GetLocation(clientIP)
	if err != nil {
		log.Printf("Error getting location: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weatherData, err := services.GetWeatherData(location.Latitude, location.Longitude)
	if err != nil {
		log.Printf("Error getting weather data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"client_ip": clientIP,
		"location":  location.City,
		"greeting":  fmt.Sprintf("Hello, %s!, the temperature is %.2f degrees Celsius in %s", visitorName, weatherData.Main.TempC, location.City),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
