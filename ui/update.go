package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	leftPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	rightPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)

	if leftPressed && !g.lastMouseLeftPressed {
		g.ProcessLeftClick()
	} else if rightPressed && !g.lastMouseRightPressed {
		g.ProcessRightClick()
	} else {
		g.ProcessNothingHappens()
	}

	g.lastMouseLeftPressed = leftPressed
	g.lastMouseRightPressed = rightPressed

	return nil
}
