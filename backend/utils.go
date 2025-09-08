package backend

func (gb *GameBackend) occupiedBy(p Point) Side {
	return gb.board[p.Y][p.X].Side

}

func (gb *GameBackend) onBoard(p Point) bool {
	if p.X < 0 || p.X >= size {
		return false
	}
	if p.Y < 0 || p.Y >= size {
		return false
	}
	return true
}

func (gb *GameBackend) isMyTurn(p Point) bool {
	return gb.turn == gb.occupiedBy(p)
}

func (gb *GameBackend) isQueen(p Point) bool {
	return gb.board[p.Y][p.X].IsQueen
}

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

func (gb *GameBackend) redCheckerMoves(p Point) []Point {
	return []Point{
		{p.X + 1, p.Y + 1},
		{p.X - 1, p.Y + 1},
	}
}

func (gb *GameBackend) blueCheckerMoves(p Point) []Point {
	return []Point{
		{p.X + 1, p.Y - 1},
		{p.X - 1, p.Y - 1},
	}
}

func (gb *GameBackend) queenPossibleMoves(p Point) []Point {
	possibleMoves := []Point{}
	for i := 1; i < size; i++ {
		p := Point{p.X + i, p.Y + i}
		if gb.onBoard(p) && gb.occupiedBy(p) == None {
			possibleMoves = append(possibleMoves, p)
		} else {
			break
		}
	}
	for i := 1; i < size; i++ {
		p := Point{p.X + i, p.Y - i}
		if gb.onBoard(p) && gb.occupiedBy(p) == None {
			possibleMoves = append(possibleMoves, p)
		} else {
			break
		}
	}
	for i := 1; i < size; i++ {
		p := Point{p.X - i, p.Y + i}
		if gb.onBoard(p) && gb.occupiedBy(p) == None {
			possibleMoves = append(possibleMoves, p)
		} else {
			break
		}
	}
	for i := 1; i < size; i++ {
		p := Point{p.X - i, p.Y - i}
		if gb.onBoard(p) && gb.occupiedBy(p) == None {
			possibleMoves = append(possibleMoves, p)
		} else {
			break
		}
	}
	return possibleMoves
}
