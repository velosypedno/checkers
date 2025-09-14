package backend

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
