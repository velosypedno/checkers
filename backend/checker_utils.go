package backend

func (gb *GameBackend) checkerAttacks(p Point) []Attack {
	attacks := []Attack{}
	attacks = append(attacks, Attack{
		Attack: Point{p.X + 1, p.Y + 1},
		Move:   Point{p.X + 2, p.Y + 2},
	})
	attacks = append(attacks, Attack{
		Attack: Point{p.X - 1, p.Y + 1},
		Move:   Point{p.X - 2, p.Y + 2},
	})
	attacks = append(attacks, Attack{
		Attack: Point{p.X + 1, p.Y - 1},
		Move:   Point{p.X + 2, p.Y - 2},
	})
	attacks = append(attacks, Attack{
		Attack: Point{p.X - 1, p.Y - 1},
		Move:   Point{p.X - 2, p.Y - 2},
	})
	return attacks
}

func (gb *GameBackend) redCheckerPossibleAttacks(p Point) []Attack {
	checkerPossibleAttacks := []Attack{}
	for _, a := range gb.checkerAttacks(p) {
		if gb.onBoard(a.Move) && gb.occupiedBy(a.Attack) == Blue && gb.occupiedBy(a.Move) == None {
			checkerPossibleAttacks = append(checkerPossibleAttacks, a)
		}
	}
	return checkerPossibleAttacks
}

func (gb *GameBackend) blueCheckerPossibleAttacks(p Point) []Attack {
	checkerPossibleAttacks := []Attack{}
	for _, a := range gb.checkerAttacks(p) {
		if gb.onBoard(a.Move) && gb.occupiedBy(a.Attack) == Red && gb.occupiedBy(a.Move) == None {
			checkerPossibleAttacks = append(checkerPossibleAttacks, a)
		}
	}
	return checkerPossibleAttacks
}

func (gb *GameBackend) redCheckerMoves(p Point) []Point {
	return []Point{
		{p.X + 1, p.Y + 1},
		{p.X - 1, p.Y + 1},
	}
}

func (gb *GameBackend) redCheckerPossibleMoves(p Point) []Point {
	possibleMoves := []Point{}
	moves := gb.redCheckerMoves(p)
	for _, m := range moves {
		if gb.onBoard(m) && gb.occupiedBy(m) == None {
			possibleMoves = append(possibleMoves, m)
		}
	}
	return possibleMoves
}

func (gb *GameBackend) blueCheckerMoves(p Point) []Point {
	return []Point{
		{p.X + 1, p.Y - 1},
		{p.X - 1, p.Y - 1},
	}
}

func (gb *GameBackend) blueCheckerPossibleMoves(p Point) []Point {
	possibleMoves := []Point{}
	moves := gb.blueCheckerMoves(p)
	for _, m := range moves {
		if gb.onBoard(m) && gb.occupiedBy(m) == None {
			possibleMoves = append(possibleMoves, m)
		}
	}
	return possibleMoves
}

func (gb *GameBackend) currentCheckerPossibleAttacks(p Point) []Attack {
	curSide := gb.occupiedBy(p)
	if curSide == Red {
		return gb.redCheckerPossibleAttacks(p)
	}
	if curSide == Blue {
		return gb.blueCheckerPossibleAttacks(p)
	}
	return []Attack{}
}

func (gb *GameBackend) canCurrentCheckerBecomeQueen(p Point) bool {
	curSide := gb.occupiedBy(p)
	if curSide == None {
		return false
	}
	if curSide == Red {
		return p.Y == size-1
	}
	return p.Y == 0
}
