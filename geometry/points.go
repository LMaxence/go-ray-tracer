package geometry

import "math"

type Point struct {
	X float64
	Y float64
	Z float64
}

func (p1 Point) Add(p2 Point) Point {
	X := p1.X + p2.X
	Y := p1.Y + p2.Y
	Z := p1.Z + p2.Z
	return Point{X: X, Y: Y, Z: Z}
}

func (p1 Point) Diff(p2 Point) Point {
	X := p1.X - p2.X
	Y := p1.Y - p2.Y
	Z := p1.Z - p2.Z
	return Point{X: X, Y: Y, Z: Z}
}

func (p1 Point) Dot(p2 Point) float64 {
	X := p1.X * p2.X
	Y := p1.Y * p2.Y
	Z := p1.Z * p2.Z
	return X + Y + Z
}

func Dist(p1 Point, p2 Point) float64 {
	dx := math.Pow(p1.X-p2.X, 2)
	dy := math.Pow(p1.Y-p2.Y, 2)
	dz := math.Pow(p1.Z-p2.Z, 2)
	return math.Sqrt(dx + dy + dz)
}

func Norm(p Point) float64 {
	return Dist(p, Point{0, 0, 0})
}

func Scale(p Point, a float64) Point {
	return Point{
		X: p.X * a,
		Y: p.Y * a,
		Z: p.Z * a,
	}
}

func Normalize(p Point) Point {
	if Norm(p) > 0 {
		return Scale(p, 1/Norm(p))
	}
	return p
}
