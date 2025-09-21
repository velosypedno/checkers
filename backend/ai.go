package backend

import (
	"math"
)

type Move struct {
	From     Point
	To       Point
	IsAttack bool
}

type MinimaxResult struct {
	Move  Move
	Score int
}

func EvaluateBoard(gb *GameBackend) int {
	score := 0

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			checker := gb.board[y][x]
			if checker.Side == None {
				continue
			}

			pieceValue := 2
			if checker.IsQueen {
				pieceValue = 8
			}
			positionBonus := 0
			if checker.Side == Red && !checker.IsQueen {
				positionBonus = y
			} else if checker.Side == Blue && !checker.IsQueen {
				positionBonus = (size - 1) - y
			}

			totalValue := pieceValue + positionBonus

			if checker.Side == Blue {
				score += totalValue
			} else {
				score -= totalValue
			}
		}
	}

	return score
}

func GetAllPossibleMoves(gb *GameBackend) []Move {
	moves := []Move{}

	if gb.IsBattlePresent() {
		attackers := gb.GetCheckersThatCanAttack()
		for _, attacker := range attackers {
			attacks := gb.PossibleAttacks(attacker.X, attacker.Y)
			for _, attack := range attacks {
				moves = append(moves, Move{
					From:     attacker,
					To:       attack.Move,
					IsAttack: true,
				})
			}
		}
	} else {
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				if gb.board[y][x].Side == gb.turn {
					possibleMoves := gb.AllowedMoves(Point{x, y})
					for _, move := range possibleMoves {
						moves = append(moves, Move{
							From:     Point{x, y},
							To:       move,
							IsAttack: false,
						})
					}
				}
			}
		}
	}

	return moves
}

func MakeMove(gb *GameBackend, move Move) *GameBackend {
	gbCopy := &GameBackend{
		board:  gb.board,
		turn:   gb.turn,
		locked: gb.locked,
	}

	if move.IsAttack {
		gbCopy.Attack(move.From, move.To)
	} else {
		gbCopy.Move(move.From, move.To)
	}

	return gbCopy
}

func Minimax(gb *GameBackend, depth int, isMaximizing bool, alpha, beta int) int {
	if depth == 0 {
		return EvaluateBoard(gb)
	}

	moves := GetAllPossibleMoves(gb)
	if len(moves) == 0 {
		return EvaluateBoard(gb)
	}

	if isMaximizing {
		maxEval := math.MinInt32
		for _, move := range moves {
			gbCopy := MakeMove(gb, move)
			var eval int
			if gbCopy.IsLocked() {
				eval = Minimax(gbCopy, depth, true, alpha, beta)
			} else {
				eval = Minimax(gbCopy, depth-1, false, alpha, beta)
			}

			maxEval = max(maxEval, eval)
			alpha = max(alpha, eval)
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		minEval := math.MaxInt32
		for _, move := range moves {
			gbCopy := MakeMove(gb, move)
			var eval int
			if gbCopy.IsLocked() {
				eval = Minimax(gbCopy, depth, false, alpha, beta)
			} else {
				eval = Minimax(gbCopy, depth-1, true, alpha, beta)
			}
			minEval = min(minEval, eval)
			beta = min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}

func GetBestMove(gb *GameBackend, depth int) Move {
	moves := GetAllPossibleMoves(gb)
	if len(moves) == 0 {
		return Move{}
	}

	bestMove := moves[0]
	bestScore := math.MinInt32

	if gb.turn == Blue {
		for _, move := range moves {
			gbCopy := MakeMove(gb, move)
			score := Minimax(gbCopy, depth-1, false, math.MinInt32, math.MaxInt32)
			if score > bestScore {
				bestScore = score
				bestMove = move
			}
		}
	} else {
		bestScore = math.MaxInt32
		for _, move := range moves {
			gbCopy := MakeMove(gb, move)
			score := Minimax(gbCopy, depth-1, true, math.MinInt32, math.MaxInt32)
			if score < bestScore {
				bestScore = score
				bestMove = move
			}
		}
	}

	return bestMove
}

func GetBestMoveForRed(game *GameBackend, depth int) Move {
	if game.turn != Red {
		return Move{}
	}

	return GetBestMove(game, depth)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
