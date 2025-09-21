package connect4

// Piece represents a player's piece and also turns.
type Piece uint
type Move uint

//Here we will define the player structure to keep track of the player information
type Player struct {
	Name      string
	TurnCount int
	Piece     Piece
	IsHuman   bool
}

//Generic Incrementer closure
func incrementer() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

const PlayerIcon Piece = 1
const CpuIcon Piece = 2
const Empty Piece = 0

func (piece Piece) opposite() Piece {
	if piece == Empty {
		return piece
	}
	return 3 - piece
}

// Description of a piece; useful to be used in the
// description of a board
func (piece Piece) String() string {
	switch piece {
	case PlayerIcon:
		return "+"
	case CpuIcon:
		return "*"
	default:
		return " "
	}
}

// // Missplaced
// // Return if move is in the given list of Moves or not
// func contains(list []Move, move Move) bool {
// 	for _, m := range list {
// 		if m == move {
// 			return true
// 		}
// 	}
// 	return false
// }
