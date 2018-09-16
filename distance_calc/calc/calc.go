package calc

import (
	p "distance_calc/point"
	"distance_calc/services"
	"net/http"
)

type DistanceResult struct {
	Distance int64
	Duration int64
}

//var origin = Point{55.8041983, 37.5831677}
//var destination = Point{55.9663444, 37.4159007}

func Calc(service services.Service, start p.Point, end p.Point) (*DistanceResult, error) {
	//result := DistanceResult{25178, 1957}
	err := service.Request(&http.Client{}, start, end)
	if err != nil {
		return nil, err
	}
	result := &DistanceResult{service.Distance(), service.Duration()}
	return result, err
}
