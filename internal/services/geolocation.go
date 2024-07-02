package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Geolocation struct {
	Status      string  `json:"status"`
	CountryName string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	RegionCode  string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	TimeZone    string  `json:"timezone"`
	IP          string  `json:"query"`
}

func GetLocation(ip string) (*Geolocation, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
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
