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

	if !gb.IsCandidateToAttack(dst) {
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
		if gb.currentFigureHasPossibleAttacks(candidate) {
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
	if !gb.onBoard(p) {
		return false
	}
	if !gb.isMyTurn(p) {
		return false
	}
	if gb.IsBattlePresent() {
		return false
	}
	return gb.currentFigureHasPossibleMoves(p)
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
		if gb.IsCandidateToAttack(candidate) {
			return true
		}
	}
	return false
}

func (gb *GameBackend) IsCandidateToAttack(p Point) bool {
	if !gb.onBoard(p) {
		return false
	}
	if !gb.isMyTurn(p) {
		return false
	}
	if !gb.currentFigureHasPossibleAttacks(p) {
		return false
	}
	return true
}

func (gb *GameBackend) IsPossibleAttack(src, dst Point) bool {
	if !gb.onBoard(src) || !gb.onBoard(dst) {
		return false
	}
	if !gb.isMyTurn(src) {
		return false
	}
	if !gb.currentFigureHasPossibleAttacks(src) {
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
