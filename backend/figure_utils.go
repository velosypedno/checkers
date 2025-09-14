package backend

func (gb *GameBackend) currentFigurePossibleMoves(p Point) []Point {
	curSide := gb.occupiedBy(p)
	if gb.isQueen(p) {
		return gb.queenPossibleMoves(p)
	}
	if curSide == Red {
		return gb.redCheckerPossibleMoves(p)
	}
	if curSide == Blue {
		return gb.blueCheckerPossibleMoves(p)
	}
	return []Point{}
}

func (gb *GameBackend) currentFigurePossibleAttacks(p Point) []Attack {
	if gb.isQueen(p) {
		return gb.currentQueenPossibleAttacks(p)
	} else {
		return gb.currentCheckerPossibleAttacks(p)
	}
}

func (gb *GameBackend) currentFigureHasPossibleAttacks(p Point) bool {
	possibleAttacks := gb.currentFigurePossibleAttacks(p)
	return len(possibleAttacks) > 0
}

func (gb *GameBackend) currentFigureHasPossibleMoves(p Point) bool {
	return len(gb.currentFigurePossibleMoves(p)) > 0
}

func (gb *GameBackend) canCurrentFigureBecomeQueen(p Point) bool {
	if gb.isQueen(p) {
		return false
	}
	return gb.canCurrentCheckerBecomeQueen(p)
}
