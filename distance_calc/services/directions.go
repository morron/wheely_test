package services

import (
	p "distance_calc/point"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ResponseRow struct {
	Value int64 `json:"value"`
}

type ResponseLegs struct {
	Distance ResponseRow `json:"distance"`
	Duration ResponseRow `json:"duration"`
}

type ResponseRoutes struct {
	Legs []ResponseLegs `json:"legs"`
}

type DistanceResponseBody struct {
	Status string           `json:"status"`
	Routes []ResponseRoutes `json:"routes"`
}

type Directions struct {
	Result *DistanceResponseBody
}

func (d *Directions) URL() string {
	return "https://maps.googleapis.com/maps/api/directions/json"
}

func (d *Directions) Distance() int64 {
	return d.Result.Routes[0].Legs[0].Distance.Value
}

func (d *Directions) Duration() int64 {
	return d.Result.Routes[0].Legs[0].Duration.Value
}

func (d *Directions) Request(client HttpClient, origin p.Point, destination p.Point) error {
	request, err := http.NewRequest("GET", d.URL(), nil)
	if err != nil {
		return err

	}

	q := request.URL.Query()
	q.Add("key", os.Getenv("APIKEY"))
	q.Add("origin", fmt.Sprint(&origin))
	q.Add("destination", fmt.Sprint(&destination))

	request.URL.RawQuery = q.Encode()
	fmt.Println(request.URL)

	httpResponse, err := client.Do(request)
	if err != nil {
		return err

	}

	var rawBody []byte
	if httpResponse.StatusCode == 200 {
		rawBody, err = ioutil.ReadAll(httpResponse.Body)

	}

	if err != nil {
		return err

	}

	err = json.Unmarshal(rawBody, &d.Result)

	return nil

}
