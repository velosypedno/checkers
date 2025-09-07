package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/velosypedno/checkers/backend"
)

func (g *Game) ProcessLeftClick() {
	x, y := ebiten.CursorPosition()
	xCell := (x - offsetXY - frameSizePX) / cellSizePX
	yCell := (y - offsetXY - frameSizePX) / cellSizePX
	p := backend.Point{X: xCell, Y: yCell}

	switch g.state {
	case Nothing:
		if g.gameBackend.CanMove(p) {
			g.ChosenToMove(p)
		}
	case ChosenToMove:
		if g.gameBackend.CanMove(p) {
			g.ChosenToMove(p)
			break
		}
		if !g.gameBackend.IsPossibleMove(*g.selected, p) {
			g.Nothing()
			break
		}
		g.gameBackend.Move(*g.selected, p)
		g.Nothing()
	case ShouldAttack:
		// Transition to ChosenToAttack
		if g.gameBackend.IsCandidateToAttack(p) {
			g.ChosenToAttack(p)
		}
	case ChosenToAttack:
		// Transition to Nothing
		if g.gameBackend.IsPossibleAttack(*g.selected, p) {
			g.gameBackend.Attack(*g.selected, p)
			g.Nothing()
			break
		}
		// Transition to ChosenToAttack
		if g.gameBackend.IsCandidateToAttack(p) {
			g.ChosenToAttack(p)
			break
		}
		// Transition to ShouldAttack
		g.ShouldAttack()

	case Locked:
		// Transition to Nothing
		if g.gameBackend.IsPossibleAttack(*g.locked, p) {
			g.gameBackend.Attack(*g.locked, p)
			g.Nothing()
		}
	}
}

func (g *Game) ProcessRightClick() {
	switch g.state {
	case ChosenToMove:
		// Transition to Nothing
		g.Nothing()
	case ChosenToAttack:
		// Transition to ShouldAttack
		g.ShouldAttack()
	case Nothing:
	case ShouldAttack:
	case Locked:
	}
}

func (g *Game) ProcessNothingHappens() {
	switch g.state {
	case Nothing:
		// Transition to Locked
		if g.gameBackend.IsLocked() {
			g.Locked()
			break
		}
		// Transition to ShouldAttack
		if g.gameBackend.IsBattlePresent() {
			g.ShouldAttack()
			break
		}
	case ShouldAttack:
	case ChosenToMove:
	case ChosenToAttack:
	case Locked:
	}
}
