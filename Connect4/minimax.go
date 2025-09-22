package connect4

import (
	"math"
)

// Find the best possible outcome evaluation for originalPlayer
// depth is initially the maximum depth
func MiniMax(b C4Board, maximizing bool, p Player, depth uint) float32 {
	//if the game is over then we shouldn't be trying to go through the recursion
	if b.IsGameOver() {
		return 0
	}
	// Base case — terminal position or maximum depth reached
	if b.IsWin() || b.IsDraw() || depth == 0 {
		return b.Evaluate(p.Piece)
	}

	// Recursive case - maximize your gains or minimize the opponent's gains
	if maximizing {
		var bestEval float32 = -math.MaxFloat32 // arbitrarily low starting point
		for _, move := range b.LegalMoves() {
			result := MiniMax(b.MakeMove(p, move), false, p, depth-1)
			if result > bestEval {
				bestEval = result
			}
		}
		return bestEval
	} else { // minimizing
		var worstEval float32 = math.MaxFloat32
		for _, move := range b.LegalMoves() {
			result := MiniMax(b.MakeMove(p, move), true, p, depth-1)
			if result < worstEval {
				worstEval = result
			}
		}
		return worstEval
	}
}

// Eval represents a move evaluation
type Eval struct {
	m Move
	f float32
}

// ConcurrentFindBestMove finds the best possible move in
// the current position looking up to depth ahead.
// This version looks at each legal move from the starting position
// concurrently (runs minimax on each legal move concurrently)
func ConcurrentFindBestMove(b C4Board, p Player, depth uint) Move {
	var bestMove Move
	var bestScore float32 = -math.MaxFloat32
	legalMoves := b.LegalMoves()

	scores := make(chan Eval, len(legalMoves))

	for _, move := range legalMoves {
		go func(move Move) {
			var e Eval
			e.m = move
			e.f = MiniMax(b.MakeMove(p, move), false, b.turn, depth)
			scores <- e
		}(move)
	}

	for i := 0; i < len(legalMoves); i++ {
		eval := <-scores
		//fmt.Printf("m: %d, f: %f\n", eval.m, eval.f)
		if eval.f > bestScore {
			bestScore = eval.f
			bestMove = eval.m
		}
	}
	close(scores)

	return bestMove
}

// FindBestMove finds the best possible move in the current position
// looking up to depth ahead
// The Function will find the best move on the provided board for the player p that is providedin paramters
func FindBestMove(b C4Board, p Player, depth uint) Move {
	var bestMove Move
	var bestScore float32 = -math.MaxFloat32

	for _, move := range b.LegalMoves() {
		if score := MiniMax(b.MakeMove(p, move), false, b.turn, depth); score > bestScore {
			bestMove = move
			bestScore = score
		}
	}

	return bestMove
}
