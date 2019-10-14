package core

import (
	"graphics/geometry"
	"image"
	"image/color"
	"image/png"
	"os"
)

var viewportSize = 1
var viewportDepth = 1

type Canvas struct {
	W        int
	H        int
	ColorMap [][]color.RGBA
}

func NewCanvas(w int, h int) Canvas {
	c := Canvas{w, h, nil}
	c.ColorMap = make([][]color.RGBA, w)
	for i := 0; i < w; i++ {
		c.ColorMap[i] = make([]color.RGBA, h)
	}
	return c
}

func (c Canvas) ToViewport(x int, y int) geometry.Point {
	widthRatio := float64(viewportSize) / float64(c.W)
	heighthRatio := float64(viewportSize) / float64(c.H)
	vx, vy := float64(x)*widthRatio, float64(y)*heighthRatio
	return geometry.Point{X: vx, Y: vy, Z: float64(viewportDepth)}
}

func (c Canvas) Render(filename string) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{c.W, c.H}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	f, _ := os.Create(filename)
	for x := 0; x < c.W; x++ {
		for y := 0; y < c.H; y++ {
			img.Set(x, c.H-y-1, c.ColorMap[x][y])
		}
	}
	png.Encode(f, img)
}
