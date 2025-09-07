package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func lighten(c color.Color, fraction float64) color.Color {
	r, g, b, a := c.RGBA()
	R := float64(r>>8) + (255-float64(r>>8))*fraction
	G := float64(g>>8) + (255-float64(g>>8))*fraction
	B := float64(b>>8) + (255-float64(b>>8))*fraction
	return color.RGBA{uint8(R), uint8(G), uint8(B), uint8(a >> 8)}
}

func drawChecker(size int, outer color.Color) *ebiten.Image {
	img := ebiten.NewImage(size, size)

	cx, cy := float64(size)/2, float64(size)/2
	outerR := float64(size) * 0.5
	innerR := outerR * 0.62
	inner := lighten(outer, 0.4)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dx := float64(x) - cx
			dy := float64(y) - cy
			dist := math.Hypot(dx, dy)

			switch {
			case dist <= innerR:
				img.Set(x, y, inner)
			case dist <= outerR:
				img.Set(x, y, outer)
			default:
				img.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}

	return img
}

func GetBlueChecker(size int) *ebiten.Image {
	return drawChecker(size, BlueCheckerColor)
}

func GetRedChecker(size int) *ebiten.Image {
	return drawChecker(size, RedCheckerColor)
}
