package main

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io"
	"net/http"
	"time"
)

var c *cache.Cache

// Planet represents the structure of the response from the Star Wars API
type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	Url            string   `json:"url"`
	Films          []string `json:"films"`
}

// fetches data from the Star Wars API and caches it
func FetchPlanetData() (*Planet, error) {
	resp, err := http.Get("https://swapi.dev/api/planets/1/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var planet Planet
	err = json.Unmarshal(body, &planet)
	if err != nil {
		return nil, err
	}

	// Caches the data with a 10-minute expiration
	c.Set("planetData", planet, 10*time.Minute)

	return &planet, nil
}

func main() {
	// Initializes the cache with a default expiration time of 10 minutes, and purge expired items every 1 minute
	c = cache.New(10*time.Minute, 1*time.Minute)

	http.HandleFunc("/planet", func(w http.ResponseWriter, r *http.Request) {
		// Try to get the data from the cache
		if planetData, found := c.Get("planetData"); found {
			fmt.Println("Serving from cache")
			json.NewEncoder(w).Encode(planetData)
			return
		}

		// If not in cache, fetch data from the API
		fmt.Println("Fetching from API")
		planetData, err := FetchPlanetData()
		if err != nil {
			http.Error(w, "Failed to fetch planet data", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(planetData)
	})

	http.ListenAndServe(":8080", nil)
}
