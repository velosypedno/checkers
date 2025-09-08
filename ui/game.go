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

func (g *Game) Nothing() {
	g.state = Nothing
	g.candidatesToAttack = []backend.Point{}
	g.possibleAttacksSelected = []backend.Attack{}
	g.possibleMovesOfSelected = []backend.Point{}
	g.locked = nil
	g.selected = nil
}

func (g *Game) ShouldAttack() {
	g.state = ShouldAttack
	g.candidatesToAttack = g.gameBackend.GetCheckersThatCanAttack()
	g.possibleAttacksSelected = []backend.Attack{}
	g.possibleMovesOfSelected = []backend.Point{}
	g.locked = nil
	g.selected = nil
}

func (g *Game) ChosenToMove(p backend.Point) {
	g.state = ChosenToMove
	g.candidatesToAttack = []backend.Point{}
	g.possibleAttacksSelected = []backend.Attack{}
	g.possibleMovesOfSelected = g.gameBackend.AllowedMoves(p)
	g.selected = &p
	g.locked = nil
}

func (g *Game) ChosenToAttack(p backend.Point) {
	g.state = ChosenToAttack
	g.candidatesToAttack = []backend.Point{}
	g.possibleAttacksSelected = g.gameBackend.PossibleAttacks(p.X, p.Y)
	g.possibleMovesOfSelected = []backend.Point{}
	g.selected = &p
	g.locked = nil
}

func (g *Game) Locked() {
	g.state = Locked
	g.locked = g.gameBackend.GetLocked()
	g.selected = g.locked
	g.candidatesToAttack = []backend.Point{}
	g.possibleMovesOfSelected = []backend.Point{}
	g.possibleAttacksSelected = g.gameBackend.PossibleAttacks(g.selected.X, g.selected.Y)
}
