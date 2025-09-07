package ui

import (
	"github.com/velosypedno/checkers/backend"
)

const (
	WindowWithPX = 700
	WindowHighPX = 700
)

type Game struct {
	gameBackend             *backend.GameBackend
	selected                *backend.Point
	possibleMovesOfSelected []backend.Point
	possibleAttacksSelected []backend.Attack
	candidatesToAttack      []backend.Point
}

func NewGame() *Game {
	return &Game{
		gameBackend:             backend.NewGameBackend(),
		possibleMovesOfSelected: []backend.Point{},
	}
}

func (g *Game) Layout(outsideWith, outsideHeight int) (int, int) {
	return WindowWithPX, WindowHighPX
}
