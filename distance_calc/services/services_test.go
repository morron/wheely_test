package services

import (
	"bytes"
	p "distance_calc/point"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

var distanceResponseBody *DistanceResponseBody

var start = p.Point{55.8041983, 37.5831677}
var end = p.Point{55.9663444, 37.4159007}

func init() {
	file, err := ioutil.ReadFile("./response.json")
	if err != nil {
		panic("response.json file is not found")
	}

	json.Unmarshal(file, &distanceResponseBody)
}

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	buf, _ := json.Marshal(distanceResponseBody)

	response := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer(buf)),
	}

	return &response, nil
}

func TestDirectionsRequest(t *testing.T) {
	mock := &http.Client{}
	service := new(Directions)
	err := service.Request(mock, start, end)
	if err != nil {
		t.Error(err)

	}
	assert.Equal(t, distanceResponseBody, service.Result)

}
