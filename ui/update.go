package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/velosypedno/checkers/backend"
)

func (g *Game) Update() error {
	leftPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	rightPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)

	if leftPressed && !g.lastMouseLeftPressed {
		g.ProcessLeftClick()
	}

	if rightPressed && !g.lastMouseRightPressed {
		g.ProcessRightClick()
	}

	g.lastMouseLeftPressed = leftPressed
	g.lastMouseRightPressed = rightPressed

	return nil
}

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

		// Should be extracted to ProcessNothingHappens
		if g.gameBackend.IsBattlePresent() {
			g.SetShouldAttackState()
		} else {
			g.SetNothingState()
		}
	case ShouldAttack:
		if g.gameBackend.IsCandidateToAttack(p) {
			g.SetChosenToAttackState(p)
		}
	case ChosenToAttack:
		if g.gameBackend.IsPossibleAttack(*g.selected, p) {
			g.gameBackend.Attack(*g.selected, p)
		} else {
			g.SetShouldAttackState()
		}

		// Should be extracted to ProcessNothingHappens
		if g.gameBackend.IsBattlePresent() {
			g.SetShouldAttackState()
		} else {
			g.SetNothingState()
		}
	}
}

func (g *Game) ProcessRightClick() {
	switch g.state {
	case Nothing:
		g.SetNothingState()
	case ChosenToMove:
		g.SetNothingState()
	case ShouldAttack:
	case ChosenToAttack:
		g.SetShouldAttackState()
	}
}

func (g *Game) SetNothingState() {
	g.state = Nothing
	g.candidatesToAttack = []backend.Point{}
	g.possibleAttacksSelected = []backend.Attack{}
	g.possibleMovesOfSelected = []backend.Point{}
	g.locked = nil
	g.selected = nil
}

func (g *Game) SetShouldAttackState() {
	g.state = ShouldAttack
	g.candidatesToAttack = g.gameBackend.GetCheckersThatCanAttack()
	g.possibleAttacksSelected = []backend.Attack{}
	g.possibleMovesOfSelected = []backend.Point{}
	g.locked = nil
	g.selected = nil
}

func (g *Game) SetChosenToMoveState(p backend.Point) {
	g.state = ChosenToMove
	g.candidatesToAttack = []backend.Point{}
	g.possibleAttacksSelected = []backend.Attack{}
	g.possibleMovesOfSelected = g.gameBackend.PossibleMoves(p.X, p.Y)
	g.selected = &p
	g.locked = nil
}

func (g *Game) SetChosenToAttackState(p backend.Point) {
	g.state = ChosenToAttack
	g.candidatesToAttack = []backend.Point{}
	g.possibleAttacksSelected = g.gameBackend.PossibleAttacks(p.X, p.Y)
	g.possibleMovesOfSelected = []backend.Point{}
	g.selected = &p
	g.locked = nil
}
