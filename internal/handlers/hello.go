package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MungaSoftwiz/location-web-server/internal/utils"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world"))
	visitorName := r.URL.Query().Get("visitor_name")
	if visitorName == "" {
		visitorName = "Guest"
	}

	clientIP := utils.GetClientIP(r)
	location := services.GetLocation(clientIP)
	weatherData := services.GetWeatherData(geoData.Lat, geoData.Lon)

	response := map[string]interface{}{
		"client_ip": clientIP,
		"location":  geoData.City,
		"greeting":  fmt.Sprintf("Hello, %s!, the temperature is %d degrees Celsius in %s", visitorName, weatherData.Main.TempC, location),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}