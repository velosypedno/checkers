package backend

type (
	Side  int
	Point struct {
		X, Y int
	}
	Attack struct {
		Attack, Move Point
	}
	Checker struct {
		Side    Side
		IsQueen bool
	}
	GameBackend struct {
		board  [8][8]Checker
		turn   Side
		locked *Point
	}
)

const (
	None Side = iota
	Red
	Blue
)

type GameState struct {
	Board [8][8]Checker
}

const (
	size     = 8
	redLine  = 3
	blueLine = 5
)

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

func (gb *GameBackend) Move(src, dst Point) {
	if !gb.IsPossibleMove(src, dst) {
		return
	}
	curSide := gb.turn
	gb.board[dst.Y][dst.X] = gb.board[src.Y][src.X]
	gb.board[src.Y][src.X] = Checker{None, false}
	gb.BecomeQueen(dst)
	gb.turn = oppSide[curSide]
}

func (gb *GameBackend) Attack(src, dst Point) {
	if !gb.IsPossibleAttack(src, dst) {
		return
	}
	possibleAttacks := gb.PossibleAttacks(src.X, src.Y)
	var chosenAttack Attack
	for _, pa := range possibleAttacks {
		if pa.Move.X == dst.X && pa.Move.Y == dst.Y {
			chosenAttack = pa
		}
	}
	curSide := gb.turn
	captured := chosenAttack.Attack
	gb.board[captured.Y][captured.X] = Checker{None, false}
	gb.board[dst.Y][dst.X] = gb.board[src.Y][src.X]
	gb.board[src.Y][src.X] = Checker{None, false}
	gb.BecomeQueen(dst)

	if !gb.IsCandidateToAttack(dst) {
		gb.turn = oppSide[curSide]
		gb.locked = nil
	} else {
		gb.locked = &dst
	}

}

func (gb *GameBackend) BecomeQueen(p Point) {
	if !gb.onBoard(p) {
		return
	}
	if !gb.isMyTurn(p) {
		return
	}
	if !gb.canCurrentFigureBecomeQueen(p) {
		return
	}
	curSide := gb.occupiedBy(p)
	gb.turn = oppSide[curSide]
	gb.locked = nil
	gb.board[p.Y][p.X].IsQueen = true
}
