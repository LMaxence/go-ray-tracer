package main

import (
	"graphics/core"
	"graphics/geometry"
	"image/color"
	"math"
)

var w, h = 1000, 1000

var sphere1 = geometry.Sphere{
	Center: geometry.Point{X: 0, Y: -1, Z: 3},
	Radius: 1,
	Color:  color.RGBA{R: 255, G: 0, B: 0, A: 0xff},
}

var sphere2 = geometry.Sphere{
	Center: geometry.Point{X: 2, Y: 0, Z: 4},
	Radius: 1,
	Color:  color.RGBA{R: 0, G: 0, B: 255, A: 0xff},
}

var sphere3 = geometry.Sphere{
	Center: geometry.Point{X: -2, Y: 0, Z: 4},
	Radius: 1,
	Color:  color.RGBA{R: 0, G: 255, B: 0, A: 0xff},
}

var s = core.Scene{
	BackgroundColor: color.RGBA{R: 255, G: 255, B: 255, A: 0xff},
	W:               w,
	H:               h,
	Objects:         []core.Element{sphere1, sphere2, sphere3},
}

func main() {
	O := geometry.Point{X: 0, Y: 0, Z: 0}
	c := core.NewCanvas(w, h)

	for x := 0; x < s.W; x++ {
		for y := 0; y < s.H; y++ {
			D := c.ToViewport(int(x-c.W/2), int(y-c.H/2))
			r := core.Ray{O: O, V: D}
			color := r.Trace(s, 1, math.Inf(1))
			c.ColorMap[x][y] = color
		}
	}
	c.Render("image.png")
}
