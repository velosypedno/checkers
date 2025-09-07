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

type checker struct {
	Side Side
}

type GameBackend struct {
	board  [8][8]checker
	turn   Side
	locked *Point
}

func NewGameBackend() *GameBackend {
	gb := &GameBackend{turn: Blue}

	for row := 0; row < redLine; row++ {
		for col := 0; col < size; col++ {
			if (row+col)%2 == 1 {
				gb.board[row][col] = checker{Side: Red}
			}
		}
	}

	for row := blueLine; row < size; row++ {
		for col := 0; col < size; col++ {
			if (row+col)%2 == 1 {
				gb.board[row][col] = checker{Side: Blue}
			}
		}
	}

	return gb
}

func (gb *GameBackend) IsPossibleOperation(src, dst Point) bool {
	return gb.IsPossibleMove(src, dst) || gb.IsPossibleAttack(src, dst)

}

func (gb *GameBackend) IsPossibleMove(src, dst Point) bool {
	if !gb.isMyTurn(src) {
		return false
	}
	if gb.IsBattlePresent() {
		return false
	}
	possibleMoves := gb.PossibleMoves(src.X, src.Y)
	for _, p := range possibleMoves {
		if p.X == dst.X && p.Y == dst.Y {
			return true
		}
	}
	return false
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

func (gb *GameBackend) candidatesToAttack(p Point) []Attack {
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

func (gb *GameBackend) canAttack(x, y int) bool {
	curSide := gb.occupiedBy(x, y)
	if curSide != gb.turn {
		return false
	}
	candidates := gb.candidatesToAttack(Point{x, y})
	for _, c := range candidates {
		if gb.isOnTheBoard(c.Move.X, c.Move.Y) &&
			gb.occupiedBy(c.Attack.X, c.Attack.Y) == oppSide[curSide] &&
			gb.occupiedBy(c.Move.X, c.Move.Y) == None {
			return true
		}
	}
	return false
}

func (gb *GameBackend) canMove(x, y int) bool {
	if !gb.isOnTheBoard(x, y) {
		return false
	}
	if !(gb.turn == gb.occupiedBy(x, y)) {
		return false
	}

	if gb.IsBattlePresent() {
		return false
	}
	currentCheckerSide := gb.occupiedBy(x, y)
	if currentCheckerSide == Red {
		if gb.isOnTheBoard(x+1, y+1) && gb.occupiedBy(x+1, y+1) == None {
			return true
		}
		if gb.isOnTheBoard(x-1, y+1) && gb.occupiedBy(x-1, y+1) == None {
			return true
		}
	}

	if currentCheckerSide == Blue {
		if gb.isOnTheBoard(x+1, y-1) && gb.occupiedBy(x+1, y-1) == None {
			return true
		}
		if gb.isOnTheBoard(x-1, y-1) && gb.occupiedBy(x-1, y-1) == None {
			return true
		}
	}
	return false
}

func (gb *GameBackend) occupiedBy(x, y int) Side {
	return gb.board[y][x].Side

}

func (gb *GameBackend) isOnTheBoard(x, y int) bool {
	if x < 0 || x >= size {
		return false
	}
	if y < 0 || y >= size {
		return false
	}
	return true
}

func (gb *GameBackend) isMyTurn(src Point) bool {
	return gb.turn == gb.occupiedBy(src.X, src.Y)
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
	return gb.canAttack(p.X, p.Y)
}
