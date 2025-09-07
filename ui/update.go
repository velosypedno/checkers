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
		g.selected = nil
		g.possibleMovesOfSelected = []backend.Point{}
		g.possibleAttacksSelected = []backend.Attack{}
	}
	if g.selected == nil {
		g.candidatesToAttack = g.gameBackend.GetCheckersThatCanAttack()
	}
	return nil
}
