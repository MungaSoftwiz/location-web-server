package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MungaSoftwiz/location-web-server/internal/utils"
	"github.com/MungaSoftwiz/location-web-server/internal/services"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	visitorName := r.URL.Query().Get("visitor_name")
	if visitorName == "" {
		visitorName = "Guest"
	}

	clientIP := utils.GetClientIP(r)
	location, _ := services.GetLocation(clientIP)
	weatherData, _ := services.GetWeatherData(location.Latitude, location.Longitude)

	response := map[string]interface{}{
		"client_ip": clientIP,
		"location":  location.City,
		"greeting":  fmt.Sprintf("Hello, %s!, the temperature is %.2f degrees Celsius in %s", visitorName, weatherData.Main.TempC, location.City),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}