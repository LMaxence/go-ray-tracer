package core

import (
	"graphics/geometry"
	"image/color"
	"math"
)

type Ray struct {
	O geometry.Point
	V geometry.Point
}

func (r Ray) pos(t float64) geometry.Point {
	X := r.O.X + t*(r.V.X-r.O.X)
	Y := r.O.Y + t*(r.V.Y-r.O.Y)
	Z := r.O.Z + t*(r.V.Z-r.O.Z)
	return geometry.Point{X: X, Y: Y, Z: Z}
}

func (r Ray) Trace(s Scene, tMin float64, tMax float64) color.RGBA {
	closestT := math.Inf(1)
	var closestElt Element

	for _, elt := range s.Objects {
		t1, t2 := elt.Intersect(r.O, r.V)
		// fmt.Printf("%v : %f, %f | ", elt, t1, t2)
		if t1 > tMin && t1 < tMax && t1 < closestT {
			closestT = t1
			closestElt = elt
		}
		if t2 > tMin && t2 < tMax && t2 < closestT {
			closestT = t2
			closestElt = elt
		}
	}
	// fmt.Print("\n")
	var col color.RGBA
	if closestElt == nil {
		col = s.BackgroundColor
	} else {
		col = closestElt.GetColor()
		p := r.pos(closestT)

		i := 0.0
		for _, light := range s.Lights {
			i += light.ComputeLight(
				p,
				closestElt.GetNormal(p),
				geometry.Scale(geometry.Normalize(r.V), -1),
				closestElt.GetSpecular(),
			)
		}

		i = math.Min(i, 1)
		col.R = uint8(float64(col.R) * i)
		col.G = uint8(float64(col.G) * i)
		col.B = uint8(float64(col.B) * i)
	}

	return col
}
