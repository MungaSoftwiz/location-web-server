package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const openWeatherMapAPIURL = "http://api.openweathermap.org/data/2.5/weather"

type WeatherData struct {
	Main struct {
		TempC float64 `json:"temp"`
	} `json:"main"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func GetWeatherData(lat, lon float64) (*WeatherData, error) {
	openWeatherMapAPIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if openWeatherMapAPIKey == "" {
		return nil, fmt.Errorf("API key for OpenWeatherMap is not set")
	}

	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s", openWeatherMapAPIURL, lat, lon, openWeatherMapAPIKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get weather data: %s", response.Status)
	}

	var weatherData WeatherData
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		return nil, err
	}

	return &weatherData, nil
}