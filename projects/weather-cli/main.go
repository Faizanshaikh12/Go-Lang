package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

const apiKey = "d79a73b7fa3459a22f281b8e5db06ecf"

type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func getWeather(city string) (*WeatherResponse, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	fmt.Println("url", url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil

}

func main() {
	city := flag.String("city", "London", "Name of the city to get the weather for")
	flag.Parse()

	if *city == "" {
		fmt.Println("City name is required")
		os.Exit(1)
	}

	fmt.Println("city", city)

	weather, err := getWeather(*city)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Weather in %s:\n", *city)
	fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temp)
	fmt.Printf("Humidity: %d%%\n", weather.Main.Humidity)
	fmt.Printf("Wind Speed: %.2f m/s\n", weather.Wind.Speed)
}
