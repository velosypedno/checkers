package backend

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
