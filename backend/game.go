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

func (gb *GameBackend) IsMoveable(x, y int) bool {
	if 0 > x || x >= size {
		return false
	}
	if 0 > y || y >= size {
		return false
	}
	if gb.board[y][x].Side != None {
		return true
	}

	return false
}
