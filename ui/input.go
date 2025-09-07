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
			g.SetChosenToMoveState(p)
		}
	case ChosenToMove:
		if !g.gameBackend.IsPossibleMove(*g.selected, p) {
			g.SetNothingState()
			break
		}
		g.gameBackend.Move(*g.selected, p)
		g.SetNothingState()
	case ShouldAttack:
		if g.gameBackend.IsCandidateToAttack(p) {
			g.SetChosenToAttackState(p)
		}
	case ChosenToAttack:
		if g.gameBackend.IsPossibleAttack(*g.selected, p) {
			g.gameBackend.Attack(*g.selected, p)
			g.SetNothingState()
		} else {
			g.SetShouldAttackState()
		}
	case Locked:
		if g.gameBackend.IsPossibleAttack(*g.locked, p) {
			g.gameBackend.Attack(*g.locked, p)
			g.SetNothingState()
		}
	}
}

func (g *Game) ProcessRightClick() {
	switch g.state {
	case ChosenToMove:
		g.SetNothingState()
	case ChosenToAttack:
		g.SetShouldAttackState()
	case Nothing:
	case ShouldAttack:
	case Locked:
	}
}

func (g *Game) ProcessNothingHappens() {
	switch g.state {
	case Nothing:
		if g.gameBackend.IsLocked() {
			g.SetLockedState()
		} else if g.gameBackend.IsBattlePresent() {
			g.SetShouldAttackState()
		} else {
			g.SetNothingState()
		}
	case ShouldAttack:
	case ChosenToMove:
	case ChosenToAttack:
	case Locked:
	}
}
