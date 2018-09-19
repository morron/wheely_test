package services

import (
	p "distance_calc/proto"
	"net/http"
)

type Service interface {
	Distance() int64
	Duration() int64
	Status() string
	URL() string
	Request(client HttpClient, origin *p.Point, destination *p.Point) error
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
