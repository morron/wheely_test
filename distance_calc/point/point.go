package point

import "fmt"

type Point struct {
	Lat  float64
	Long float64
}

func (p *Point) String() string {
	return fmt.Sprintf("%v,%v", p.Lat, p.Long)

}
