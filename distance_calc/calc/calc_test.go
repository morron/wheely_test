package calc

import (
	p "distance_calc/proto"
	s "distance_calc/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

var start = &p.Point{Lat: 55.8041983, Lng: 37.5831677}
var end = &p.Point{Lat: 55.9663444, Lng: 37.4159007}

func TestCalc(t *testing.T) {
	expect := &p.Distance{Distance: 25178, Duration: 1957}
	service := new(s.Directions)

	got, err := Calc(service, start, end)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expect, got)
}
