package backend

var oppSide = map[Side]Side{
	Blue: Red,
	Red:  Blue,
}

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
