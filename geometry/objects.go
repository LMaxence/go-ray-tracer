package geometry

import (
	"image/color"
	"math"
)

type Sphere struct {
	Center Point
	Radius float64
	Color  color.RGBA
}

func (s Sphere) Intersect(o Point, d Point) (float64, float64) {
	oc := o.Diff(s.Center)
	k1 := d.Dot(d)
	k2 := 2 * oc.Dot(d)
	k3 := oc.Dot(oc) - math.Pow(s.Radius, 2)

	discriminant := k2*k2 - 4*k1*k3
	if discriminant < 0 {
		return math.Inf(1), math.Inf(1)
	}

	t1 := (-k2 + math.Sqrt(discriminant)) / (2 * k1)
	t2 := (-k2 - math.Sqrt(discriminant)) / (2 * k1)
	return t1, t2
}

func (s Sphere) GetColor() color.RGBA {
	return s.Color
}
