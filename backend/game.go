package backend

type Side int

const (
	None Side = iota
	Red
	Blue
)

const (
	size     = 8
	redLine  = 3
	blueLine = 5
)

type checker struct {
	Side Side
}

type Point struct {
	X, Y int
}

type GameBackend struct {
	board [8][8]checker
}

func NewGameBackend() *GameBackend {
	gb := &GameBackend{}

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

type GameState struct {
	Board [8][8]Side
}

func (gb *GameBackend) GetState() GameState {
	state := GameState{}
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			state.Board[row][col] = gb.board[row][col].Side
		}
	}
	return state
}

func (gb *GameBackend) Move(src, dst Point) {
	if !gb.IsPossibleMove(src, dst) {
		return
	}
	gb.board[dst.Y][dst.X] = gb.board[src.Y][src.X]
	gb.board[src.Y][src.X] = checker{None}
}

func (gb *GameBackend) IsPossibleMove(src, dst Point) bool {
	possibleMoves := gb.PossibleMoves(src.X, src.Y)
	for _, p := range possibleMoves {
		if p.X == dst.X && p.Y == dst.Y {
			return true
		}
	}
	return false
}

func (gb *GameBackend) IsMoveable(x, y int) bool {
	if !gb.isOnTheBoard(x, y) {
		return false
	}
	if gb.OccupiedBy(x, y) == None {
		return false
	}
	currentCheckerSide := gb.OccupiedBy(x, y)
	if currentCheckerSide == Red {
		if gb.isOnTheBoard(x+1, y+1) && gb.OccupiedBy(x+1, y+1) == None {
			return true
		}
		if gb.isOnTheBoard(x-1, y+1) && gb.OccupiedBy(x-1, y+1) == None {
			return true
		}

		if gb.isOnTheBoard(x+1, y+1) && gb.OccupiedBy(x+1, y+1) == Blue {
			if gb.isOnTheBoard(x+2, y+2) && gb.OccupiedBy(x+2, y+2) == None {
				return true
			}
		}
		if gb.isOnTheBoard(x-1, y+1) && gb.OccupiedBy(x-1, y+1) == Blue {
			if gb.isOnTheBoard(x-2, y+2) && gb.OccupiedBy(x-2, y+2) == None {
				return true
			}
		}
	}

	if currentCheckerSide == Blue {
		if gb.isOnTheBoard(x+1, y-1) && gb.OccupiedBy(x+1, y-1) == None {
			return true
		}
		if gb.isOnTheBoard(x-1, y-1) && gb.OccupiedBy(x-1, y-1) == None {
			return true
		}

		if gb.isOnTheBoard(x+1, y-1) && gb.OccupiedBy(x+1, y-1) == Red {
			if gb.isOnTheBoard(x+2, y-2) && gb.OccupiedBy(x+2, y-2) == None {
				return true
			}
		}
		if gb.isOnTheBoard(x-1, y-1) && gb.OccupiedBy(x-1, y-1) == Red {
			if gb.isOnTheBoard(x-2, y-2) && gb.OccupiedBy(x-2, y-2) == None {
				return true
			}
		}
	}

	return false
}

func (gb *GameBackend) OccupiedBy(x, y int) Side {
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

func (gb *GameBackend) PossibleMoves(x, y int) []Point {
	possibleMoves := []Point{}
	if !gb.IsMoveable(x, y) {
		return possibleMoves
	}
	currentCheckerSide := gb.OccupiedBy(x, y)
	var possibleX int
	var possibleY int
	if currentCheckerSide == Red {
		possibleX = x + 1
		possibleY = y + 1
		if gb.isOnTheBoard(possibleX, possibleY) && gb.OccupiedBy(possibleX, possibleY) == None {
			possibleMoves = append(possibleMoves, Point{possibleX, possibleY})
		}
		possibleX = x - 1
		possibleY = y + 1
		if gb.isOnTheBoard(possibleX, possibleY) && gb.OccupiedBy(possibleX, possibleY) == None {
			possibleMoves = append(possibleMoves, Point{possibleX, possibleY})
		}
	}

	if currentCheckerSide == Blue {
		possibleX = x + 1
		possibleY = y - 1
		if gb.isOnTheBoard(possibleX, possibleY) && gb.OccupiedBy(possibleX, possibleY) == None {
			possibleMoves = append(possibleMoves, Point{possibleX, possibleY})
		}
		possibleX = x - 1
		possibleY = y - 1
		if gb.isOnTheBoard(possibleX, possibleY) && gb.OccupiedBy(possibleX, possibleY) == None {
			possibleMoves = append(possibleMoves, Point{possibleX, possibleY})
		}
	}

	return possibleMoves
}
