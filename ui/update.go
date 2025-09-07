package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/velosypedno/checkers/backend"
)

func (g *Game) Update() error {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// Step 1: Calculate position
		x, y := ebiten.CursorPosition()
		xCell := (x - offsetXY - frameSizePX) / cellSizePX
		yCell := (y - offsetXY - frameSizePX) / cellSizePX

		// Step 2: Check moves and attacks if a checker is already selected
		if g.selected != nil {
			src := backend.Point{
				X: g.selected.X,
				Y: g.selected.Y,
			}
			dst := backend.Point{
				X: xCell,
				Y: yCell,
			}
			if g.gameBackend.IsPossibleMove(src, dst) {
				g.gameBackend.Move(src, dst)
				g.selected = nil
				g.possibleMovesOfSelected = []backend.Point{}
				g.possibleAttacksSelected = []backend.Attack{}
			} else if g.gameBackend.IsPossibleAttack(src, dst) {
				g.gameBackend.Attack(src, dst)
				g.selected = nil
				g.possibleMovesOfSelected = []backend.Point{}
				g.possibleAttacksSelected = []backend.Attack{}
			}
		}

		// Step 3: Select checker if possible or reboot highlighting
		if g.gameBackend.CanBeHighlighted(xCell, yCell) {
			g.selected = &backend.Point{X: xCell, Y: yCell}
			g.possibleMovesOfSelected = g.gameBackend.PossibleMoves(xCell, yCell)
			g.possibleAttacksSelected = g.gameBackend.PossibleAttacks(xCell, yCell)
			g.candidatesToAttack = []backend.Point{}
		} else {
			g.selected = nil
			g.possibleMovesOfSelected = []backend.Point{}
			g.possibleAttacksSelected = []backend.Attack{}
			g.candidatesToAttack = g.gameBackend.GetCheckersThatCanAttack()
		}

	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.ProcessRightClick()
	}
	if g.selected == nil {
		g.candidatesToAttack = g.gameBackend.GetCheckersThatCanAttack()
	}
	return nil
}

func (g *Game) ProcessLeftClick() {
	x, y := ebiten.CursorPosition()
	xCell := (x - offsetXY - frameSizePX) / cellSizePX
	yCell := (y - offsetXY - frameSizePX) / cellSizePX

	switch g.state {
	case Nothing:
		if g.gameBackend.CanBeHighlighted(xCell, yCell) {

		}
	case Selected:
		g.SetNothingState()
	case Locked:
	case ShouldAttack:
	case CandidateChosen:
		g.SetShouldAttack()
	}
}

func (g *Game) ProcessRightClick() {
	switch g.state {
	case Nothing:
		g.SetNothingState()
	case Selected:
		g.SetNothingState()
	case Locked:
	case ShouldAttack:
	case CandidateChosen:
		g.SetShouldAttack()
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

func (g *Game) SetShouldAttack() {
	g.state = ShouldAttack
	g.candidatesToAttack = g.gameBackend.GetCheckersThatCanAttack()
	g.possibleAttacksSelected = []backend.Attack{}
	g.possibleMovesOfSelected = []backend.Point{}
	g.locked = nil
	g.selected = nil
}
