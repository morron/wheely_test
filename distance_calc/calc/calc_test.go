package calc

import (
	p "distance_calc/point"
	s "distance_calc/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

var start = p.Point{55.8041983, 37.5831677}
var end = p.Point{55.9663444, 37.4159007}

func TestCalc(t *testing.T) {
	expect := &DistanceResult{25178, 1957}
	service := new(s.Directions)

	got, err := Calc(service, start, end)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expect, got)
}
