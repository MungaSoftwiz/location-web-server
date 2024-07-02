package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MungaSoftwiz/location-web-server/internal/services"
	"github.com/MungaSoftwiz/location-web-server/internal/utils"
)

type Response struct {
	ClientIP string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

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

	tempCelsius := weatherData.Main.TempC - 273.15

	response := Response{
		ClientIP: clientIP,
		Location: location.City,
		Greeting: fmt.Sprintf("Hello, %s!, the temperature is %.2f degrees Celsius in %s", visitorName, tempCelsius, location.City),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
