package calc

import (
	p "distance_calc/proto"
	"distance_calc/services"
	"net/http"
)

//var origin = Point{55.8041983, 37.5831677}
//var destination = Point{55.9663444, 37.4159007}

func Calc(service services.Service, start *p.Point, end *p.Point) (*p.Distance, error) {
	err := service.Request(&http.Client{}, start, end)
	if err != nil {
		return nil, err
	}
	result := &p.Distance{Distance: service.Distance(), Duration: service.Duration()}
	return result, err
}
