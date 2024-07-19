package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bytedance/sonic"
)

var (
	baseEndpoint = "https://api.weather.gov/points/%s,%s" //lat,lon
)

type Stations struct {
	Context             map[string]Context `json:"@context"`
	Type                string
	Features            []Station
	ObservationStations []string
	Pagination          interface{}
}

type Station struct {
	ID         string `json:"id"`
	Type       string
	Geometry   Geometry
	Properties Properties
}

type Context struct {
	Version          string `json:"@version"`
	Wx               string
	S                string
	Geo              string
	Unit             string
	Vocab            string `json:"@vocab"`
	Geometry         interface{}
	City             string
	State            string
	Distance         interface{}
	Bearing          interface{}
	Value            interface{}
	UnitCode         interface{}
	ForecastOffice   interface{}
	ForecastGridData interface{}
	PublicZone       interface{}
	County           interface{}
}

type Geometry struct {
	Type        string
	Coordinates []float64
}

type Properties struct {
	ID                string `json:"@id"`
	Type              string `json:"@type"`
	Elevation         Elevation
	StationIdentifier string
	Name              string
	TimeZone          string
	Forecast          string
	County            string
	FireWeatherZone   string
}

type Elevation struct {
	UnitCode string
	Value    float64
}

func check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// not functional
func main() {
	latLon := []string{"30.269794", "-97.773486"}
	endpoint := fmt.Sprintf(baseEndpoint, latLon[0], latLon[1])
	req, err := http.Get(endpoint)
	check(err)
	var stations Stations
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	check(err)
	err = sonic.Unmarshal(body, &stations)
	check(err)
}
