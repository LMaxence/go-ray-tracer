package effects

import (
	"graphics/geometry"
	"math"
)

type AmbientLight struct {
	Intensity float64
}

func (l AmbientLight) ComputeLight(p geometry.Point, n geometry.Point) float64 {
	return l.Intensity
}

type PointLight struct {
	Intensity float64
	Center    geometry.Point
}

func (l PointLight) ComputeLight(p geometry.Point, n geometry.Point, v geometry.Point, s float64) float64 {
	// diffuse light
	lightDirection := geometry.Normalize(l.Center.Diff(p))
	nDirection := geometry.Normalize(n)
	factor := math.Max(0, nDirection.Dot(lightDirection))

	// specular light
	if s != -1 {
		r := geometry.Scale(nDirection, 2*nDirection.Dot(lightDirection)).Diff(lightDirection)
		specularFactor := r.Dot(v)
		if specularFactor > 0 {
			factor += l.Intensity * math.Pow(specularFactor, s)
		}
	}
	return factor
}
