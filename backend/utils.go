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

func (gb *GameBackend) currentQueenPossibleAttacks(p Point) []Attack {
	sides := []Point{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	possibleAttacks := []Attack{}
	curSide := gb.occupiedBy(p)
	var opponentFigure *Point = nil
	for _, s := range sides {
		for i := 1; i < size; i++ {
			curP := Point{p.X + (i * s.X), p.Y + (i * s.Y)}
			if !gb.onBoard(curP) {
				break
			}
			if gb.occupiedBy(curP) == oppSide[curSide] {
				if opponentFigure != nil {
					break
				} else {
					opponentFigure = &curP
					continue
				}
			}
			if gb.occupiedBy(curP) == None {
				if opponentFigure != nil {
					possibleAttacks = append(possibleAttacks, Attack{*opponentFigure, curP})
				}
			} else {
				break
			}
		}
		opponentFigure = nil
	}
	return possibleAttacks
}

func (gb *GameBackend) currentFigurePossibleAttacks(p Point) []Attack {
	if gb.isQueen(p) {
		return gb.currentQueenPossibleAttacks(p)
	} else {
		return gb.currentCheckerPossibleAttacks(p)
	}
}
