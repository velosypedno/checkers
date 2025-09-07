package backend

type GameState struct {
	Board [8][8]Side
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
	if gb.canAttack(x, y) {
		return possibleMoves
	}
	if !gb.isOnTheBoard(x, y) {
		return possibleMoves
	}
	if gb.IsBattlePresent() {
		return possibleMoves
	}
	if gb.occupiedBy(x, y) == None {
		return possibleMoves
	}

	// Check for possible moves
	curSide := gb.occupiedBy(x, y)
	var possibleX int
	var possibleY int
	if curSide == Red {
		possibleX = x + 1
		possibleY = y + 1
		if gb.isOnTheBoard(possibleX, possibleY) && gb.occupiedBy(possibleX, possibleY) == None {
			possibleMoves = append(possibleMoves, Point{possibleX, possibleY})
		}
		possibleX = x - 1
		possibleY = y + 1
		if gb.isOnTheBoard(possibleX, possibleY) && gb.occupiedBy(possibleX, possibleY) == None {
			possibleMoves = append(possibleMoves, Point{possibleX, possibleY})
		}
	}

	if curSide == Blue {
		possibleX = x + 1
		possibleY = y - 1
		if gb.isOnTheBoard(possibleX, possibleY) && gb.occupiedBy(possibleX, possibleY) == None {
			possibleMoves = append(possibleMoves, Point{possibleX, possibleY})
		}
		possibleX = x - 1
		possibleY = y - 1
		if gb.isOnTheBoard(possibleX, possibleY) && gb.occupiedBy(possibleX, possibleY) == None {
			possibleMoves = append(possibleMoves, Point{possibleX, possibleY})
		}
	}
	return possibleMoves
}

func (gb *GameBackend) PossibleAttacks(x, y int) []Attack {
	curSide := gb.occupiedBy(x, y)
	attacks := []Attack{}
	candidates := gb.candidatesToAttack(Point{x, y})
	for _, c := range candidates {
		if gb.isOnTheBoard(c.Move.X, c.Move.Y) &&
			gb.occupiedBy(c.Attack.X, c.Attack.Y) == oppSide[curSide] &&
			gb.occupiedBy(c.Move.X, c.Move.Y) == None {
			attacks = append(attacks, c)
		}
	}
	return attacks
}

func (gb *GameBackend) CanBeHighlighted(x, y int) bool {
	if !gb.isOnTheBoard(x, y) {
		return false
	}

	if !gb.isMyTurn(Point{x, y}) {
		return false
	}

	return gb.canMove(x, y) || gb.canAttack(x, y)
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
	gb.turn = oppSide[gb.turn]
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
	captured := chosenAttack.Attack
	gb.board[captured.Y][captured.X] = checker{None}
	gb.board[dst.Y][dst.X] = gb.board[src.Y][src.X]
	gb.board[src.Y][src.X] = checker{None}
	if !gb.canAttack(dst.X, dst.Y) {
		gb.turn = oppSide[gb.turn]
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
