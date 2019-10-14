package core

import (
	"graphics/geometry"
	"image/color"
)

type Element interface {
	Intersect(o geometry.Point, d geometry.Point) (float64, float64)
	GetColor() color.RGBA
}

type Scene struct {
	Objects         []Element
	BackgroundColor color.RGBA
	W               int
	H               int
}
