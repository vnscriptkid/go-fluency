package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type WeatherAPIResponse1 struct {
	Temperature float64 `json:"temperature"`
}

type WeatherAPIResponse2 struct {
	Current struct {
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}

func main() {
	// Define the list of weather API endpoints
	apiEndpoints := []string{
		"https://api.weatherapi1.com/temp",
		"https://api.weatherapi2.com/temp",
	}

	// Use scatterGather to fetch the temperatures and calculate the average
	averageTemperature := scatterGather(apiEndpoints)

	fmt.Printf("The average temperature is: %.2f\n", averageTemperature)
}

func scatterGather(endpoints []string) float64 {
	var wg sync.WaitGroup
	temperatureSum := 0.0
	totalResponses := 0

	wg.Add(len(endpoints))

	for _, endpoint := range endpoints {
		go func(api string) {
			defer wg.Done()

			temp, err := fetchTemperature(api)
			if err != nil {
				fmt.Printf("Error fetching temperature from %s: %v\n", api, err)
				return
			}

			temperatureSum += temp
			totalResponses++
		}(endpoint)
	}

	wg.Wait()

	if totalResponses == 0 {
		fmt.Println("No successful temperature responses.")
		return 0
	}

	return temperatureSum / float64(totalResponses)
}

func fetchTemperature(api string) (float64, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(api)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Normalize the temperature data based on the API's response structure
	switch api {
	case "https://api.weatherapi1.com/temp":
		var weatherResp WeatherAPIResponse1
		if err := json.Unmarshal(body, &weatherResp); err != nil {
			return 0, err
		}
		return weatherResp.Temperature, nil

	case "https://api.weatherapi2.com/temp":
		var weatherResp WeatherAPIResponse2
		if err := json.Unmarshal(body, &weatherResp); err != nil {
			return 0, err
		}
		return weatherResp.Current.TempF, nil

	default:
		return 0, fmt.Errorf("unsupported API: %s", api)
	}
}
