package core

import (
	"graphics/geometry"
	"image/color"
)

type Element interface {
	Intersect(geometry.Point, geometry.Point) (float64, float64)
	GetColor() color.RGBA
	GetNormal(geometry.Point) geometry.Point
	GetSpecular() float64
}

type Light interface {
	ComputeLight(geometry.Point, geometry.Point, geometry.Point, float64) float64
}

type Scene struct {
	Objects         []Element
	BackgroundColor color.RGBA
	Lights          []Light
	W               int
	H               int
}
