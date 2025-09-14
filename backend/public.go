package backend

type GameState struct {
	Board [8][8]Checker
}

type Point struct {
	X, Y int
}

type Attack struct {
	Attack, Move Point
}

type Side int

const (
	None Side = iota
	Red
	Blue
)

func (gb *GameBackend) AllowedMoves(p Point) []Point {
	if !gb.onBoard(p) {
		return []Point{}
	}
	if gb.occupiedBy(p) == None {
		return []Point{}
	}
	if !gb.isMyTurn(p) {
		return []Point{}
	}
	if gb.IsBattlePresent() {
		return []Point{}
	}
	return gb.currentFigurePossibleMoves(p)
}

func (gb *GameBackend) PossibleAttacks(x, y int) []Attack {
	p := Point{x, y}
	return gb.currentFigurePossibleAttacks(p)
}

func (gb *GameBackend) GetState() GameState {
	state := GameState{}
	state.Board = gb.board
	return state
}

func (gb *GameBackend) Move(src, dst Point) {
	if !gb.IsPossibleMove(src, dst) {
		return
	}
	curSide := gb.turn
	gb.board[dst.Y][dst.X] = gb.board[src.Y][src.X]
	gb.board[src.Y][src.X] = Checker{None, false}
	gb.tryToBecameQueen(dst)
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
	gb.tryToBecameQueen(dst)

	if !gb.canAttack(dst.X, dst.Y) {
		gb.turn = oppSide[curSide]
		gb.locked = nil
	} else {
		gb.locked = &dst
	}

}

func (gb *GameBackend) GetCheckersThatCanAttack() []Point {
	candidates := []Point{}
	for x := range size {
		for y := range size {
			if gb.board[y][x].Side == gb.turn {
				candidates = append(candidates, Point{x, y})
			}
		}
	}
	checkersThatCanAttack := []Point{}
	for _, candidate := range candidates {
		if gb.canAttack(candidate.X, candidate.Y) {
			checkersThatCanAttack = append(checkersThatCanAttack, candidate)
		}
	}
	return checkersThatCanAttack
}

func (gb *GameBackend) GetLocked() *Point {
	return gb.locked
}

func (gb *GameBackend) IsLocked() bool {
	return gb.GetLocked() != nil
}

func (gb *GameBackend) CanMove(p Point) bool {
	return gb.canMove(p)
}
