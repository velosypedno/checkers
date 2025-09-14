package backend

const (
	size     = 8
	redLine  = 3
	blueLine = 5
)

var oppSide = map[Side]Side{
	Blue: Red,
	Red:  Blue,
}

type Checker struct {
	Side    Side
	IsQueen bool
}

type GameBackend struct {
	board  [8][8]Checker
	turn   Side
	locked *Point
}

func NewGameBackend() *GameBackend {
	gb := &GameBackend{turn: Blue}

	for row := 0; row < redLine; row++ {
		for col := 0; col < size; col++ {
			if (row+col)%2 == 1 {
				gb.board[row][col] = Checker{Side: Red}
			}
		}
	}

	for row := blueLine; row < size; row++ {
		for col := 0; col < size; col++ {
			if (row+col)%2 == 1 {
				gb.board[row][col] = Checker{Side: Blue}
			}
		}
	}

	return gb
}

func (gb *GameBackend) IsPossibleMove(src, dst Point) bool {
	if !gb.onBoard(src) || !gb.onBoard(dst) {
		return false
	}
	if !gb.isMyTurn(src) {
		return false
	}
	if gb.IsBattlePresent() {
		return false
	}
	allowedMoves := gb.AllowedMoves(src)
	for _, p := range allowedMoves {
		if p.X == dst.X && p.Y == dst.Y {
			return true
		}
	}
	return false
}

func (gb *GameBackend) tryToBecameQueen(p Point) {
	curSide := gb.occupiedBy(p)
	if curSide == None {
		return
	}
	if curSide == Red {
		if p.Y == size-1 {
			gb.board[p.Y][p.X].IsQueen = true
			gb.turn = oppSide[curSide]
			gb.locked = nil
		}

	}

	if curSide == Blue {
		if p.Y == 0 {
			gb.board[p.Y][p.X].IsQueen = true
			gb.turn = oppSide[curSide]
			gb.locked = nil
		}
	}
}

func (gb *GameBackend) IsPossibleAttack(src, dst Point) bool {
	if !gb.isMyTurn(src) {
		return false
	}
	possibleAttacks := gb.PossibleAttacks(src.X, src.Y)
	for _, p := range possibleAttacks {
		if p.Move.X == dst.X && p.Move.Y == dst.Y {
			return true
		}
	}
	return false
}

func (gb *GameBackend) canAttack(x, y int) bool {
	curSide := gb.occupiedBy(Point{x, y})
	if curSide != gb.turn {
		return false
	}
	possibleAttacks := gb.currentFigurePossibleAttacks(Point{x, y})
	return len(possibleAttacks) > 0
}

func (gb *GameBackend) canMove(p Point) bool {
	return len(gb.AllowedMoves(p)) > 0
}

func (gb *GameBackend) IsBattlePresent() bool {
	candidates := []Point{}
	for x := range size {
		for y := range size {
			if gb.board[y][x].Side == gb.turn {
				candidates = append(candidates, Point{x, y})
			}
		}
	}
	for _, candidate := range candidates {
		if gb.canAttack(candidate.X, candidate.Y) {
			return true
		}
	}
	return false
}

func (gb *GameBackend) IsCandidateToAttack(p Point) bool {
	return gb.onBoard(p) && gb.canAttack(p.X, p.Y)
}
