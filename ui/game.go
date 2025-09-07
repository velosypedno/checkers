package ui

import (
	"github.com/velosypedno/checkers/backend"
)

const (
	WindowWithPX = 700
	WindowHighPX = 700
)

type GameState int

const (
	Nothing = iota
	ChosenToMove
	Locked
	ShouldAttack
	ChosenToAttack
)

type Game struct {
	gameBackend             *backend.GameBackend
	selected                *backend.Point
	possibleMovesOfSelected []backend.Point
	possibleAttacksSelected []backend.Attack
	candidatesToAttack      []backend.Point
	locked                  *backend.Point
	state                   GameState

	lastMouseLeftPressed  bool
	lastMouseRightPressed bool
}

func NewGame() *Game {
	return &Game{
		gameBackend:             backend.NewGameBackend(),
		possibleMovesOfSelected: []backend.Point{},
		possibleAttacksSelected: []backend.Attack{},
		selected:                nil,
		locked:                  nil,
		candidatesToAttack:      []backend.Point{},
		state:                   Nothing,
	}
}

func (g *Game) Layout(outsideWith, outsideHeight int) (int, int) {
	return WindowWithPX, WindowHighPX
}
