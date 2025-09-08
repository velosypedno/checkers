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

func (gb *GameBackend) PossibleMoves(x, y int) []Point {
	// Preconditions
	possibleMoves := []Point{}
	p := Point{x, y}
	if !gb.onBoard(p) {
		return possibleMoves
	}
	if !gb.isMyTurn(p) {
		return possibleMoves
	}
	if gb.IsBattlePresent() {
		return possibleMoves
	}
	if gb.occupiedBy(p) == None {
		return possibleMoves
	}
	var candidates []Point
	curSide := gb.occupiedBy(p)
	if gb.isQueen(p) {
		return append(possibleMoves, gb.queenPossibleMoves(p)...)
	}
	if curSide == Red {
		candidates = gb.redCheckerMoves(p)
	}
	if curSide == Blue {
		candidates = gb.blueCheckerMoves(p)
	}

	for _, c := range candidates {
		if gb.onBoard(c) && gb.occupiedBy(c) == None {
			possibleMoves = append(possibleMoves, c)
		}
	}
	return possibleMoves
}

func (gb *GameBackend) PossibleAttacks(x, y int) []Attack {
	p := Point{x, y}
	curSide := gb.occupiedBy(p)
	attacks := []Attack{}
	candidates := gb.checkerAttacks(p)
	for _, c := range candidates {
		if gb.onBoard(c.Move) &&
			gb.occupiedBy(c.Attack) == oppSide[curSide] &&
			gb.occupiedBy(c.Move) == None {
			attacks = append(attacks, c)
		}
	}
	return attacks
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
	gb.TryToBecameQueen(dst)
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
	gb.TryToBecameQueen(dst)

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
	return gb.canMove(p.X, p.Y)
}
