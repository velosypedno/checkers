package main

import (
	"image/color"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWith  = 700
	windowHigh  = 700
	title       = "Checkers"
	boardSizePX = 640
	offsetXY    = 20
	frameSizePX = 5
	cellSizePX  = boardSizePX / boardSize

	boardSize = 8
)

var (
	backgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func DrawBoard(boardImg *ebiten.Image) {
	boardImg.Fill(frameColor)
	isBlack := true
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			cellImg := ebiten.NewImage(cellSizePX, cellSizePX)
			if isBlack = !isBlack; isBlack {
				cellImg.Fill(color.Black)
			} else {
				cellImg.Fill(color.White)
			}
			op := &ebiten.DrawImageOptions{}
			xOffset := frameSizePX + i*cellSizePX
			yOffset := frameSizePX + j*cellSizePX
			op.GeoM.Translate(float64(xOffset), float64(yOffset))
			boardImg.DrawImage(cellImg, op)
		}
		isBlack = !isBlack
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	boardImg := ebiten.NewImage(boardSizePX+frameSizePX*2, boardSizePX+frameSizePX*2)
	DrawBoard(boardImg)
	op := &ebiten.DrawImageOptions{}
	xOffset := (windowWith - boardSizePX - frameSizePX*2) / 2
	yOffset := (windowHigh - boardSizePX - frameSizePX*2) / 2
	op.GeoM.Translate(float64(xOffset), float64(yOffset))
	screen.DrawImage(boardImg, op)
}

func (g *Game) Layout(outsideWith, outsideHeight int) (int, int) {
	return windowWith, windowHigh
}

func main() {
	ebiten.SetWindowSize(windowWith, windowHigh)
	ebiten.SetWindowTitle(title)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
