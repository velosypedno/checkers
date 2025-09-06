package main

import (
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/velosypedno/checkers/ui"
)

func main() {
	ebiten.SetWindowSize(ui.WindowWithPX, ui.WindowHighPX)
	ebiten.SetWindowTitle(ui.Title)
	if err := ebiten.RunGame(ui.NewGame()); err != nil {
		log.Fatal(err)
	}
}
