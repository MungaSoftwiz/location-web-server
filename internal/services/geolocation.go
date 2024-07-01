package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Geolocation struct {
	IP            string  `json:"ip"`
	CountryName   string  `json:"country_name"`
	RegionName    string  `json:"region_name"`
	City          string  `json:"city"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

func GetLocation(ip string) (*Geolocation, error) {
	url := fmt.Sprintf("https://api.ipbase.com/v1/json/%s", ip)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch geolocation: %s", response.Status)
	}

	var geo Geolocation
	if err := json.NewDecoder(response.Body).Decode(&geo); err != nil {
		return nil, err
	}

	return &geo, nil
}
