package connect4

import (
	"fmt"
	"strings"
)

// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// Structures and basic board functions
// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
const (
	NumRows = 6
	NumCols = 7
)

//------------------------------------------------------

// C4Board holds the board state. Use NumCols / NumRows for array sizes.
type C4Board struct {
	numRows  uint
	numCols  uint
	position [NumCols][NumRows]Piece // position[col][row]
	colCount [NumCols]uint           // how many pieces are in a given column
	turn     Player                  // who's turn it is to play
}

// Segment is a contiguous four-piece slice used for scoring/checking wins
type Segment [4]Piece

// NewBoard returns an initialized Connect4 board
func NewBoard() C4Board {
	b := C4Board{
		numRows: NumRows,
		numCols: NumCols,
		turn:    Player{Piece: PlayerIcon},
	}
	return b
}

// Adjusting the turn with the player piece
func (b C4Board) adjustTurn(player Player) C4Board {
	b.turn = player
	return b
}

func (b C4Board) String() string {
	var sb strings.Builder

	if b.numCols == 0 || b.numRows == 0 {
		return "Empty Board"
	}

	for i := int(b.numRows) - 1; i >= 0; i-- {
		sb.WriteString("|")
		for j := 0; j < int(b.numCols); j++ {
			sb.WriteString(b.position[j][i].String())
			sb.WriteString("|")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

//------------------------------------------------------
//------------------------------------------------------
//------------------------------------------------------
//------------------------------------------------------
//------------------------------------------------------
//Player interactions with the board
//------------------------------------------------------
//------------------------------------------------------
//------------------------------------------------------
//------------------------------------------------------
//------------------------------------------------------

// GEneric Make Move Functions that is utilized without a column as this will prompt for a column
// This will then check that the entered in value is a legal move
func (board C4Board) MakePlayerMove(p Player) C4Board {
	//If we have no more moves to make because the game is over then whyy do anything
	// This is also a failsafe for the recursion to prevent an infinite loop
	if board.IsGameOver() {
		fmt.Println("The game is over, unexpected MakeMove call...")
		fmt.Println("Proceeding without making a move...")
		return board
	}

	var col Move
	fmt.Println("Enter a Column you would like to insert in(0-6): ")
	for {
		//Going to scan the input, if it is a legal column and is not an error then we want to
		//make to move and then return the new board
		if _, err := fmt.Scanln(&col); err == nil && (col <= NumCols || col >= 0) && board.determineIfLegalMove(col) {
			return board.MakeMove(p, col)
		} else {
			fmt.Println("That was not a legal move, please try again: ")
			return board.MakePlayerMove(p) //Recursively call the function until a legal move is entered
		}

		return board.MakeMove(p, col)
	}
}

// Calulcate that the column provided that was enterered was a legal move
func (board C4Board) determineIfLegalMove(col Move) bool {
	for _, value := range board.LegalMoves() {
		if value == col {
			return true
		}
	}
	return false
}

// MakeMove puts a piece in column col.
// Returns a copy of the board with the move made.
// Does not check if the column is full (assumes legal move).
func (board C4Board) MakeMove(p Player, col Move) C4Board {
	b := board
	piece := p.Piece

	// board.colCount[col] will be the empty space in the column
	// technically this can error however it shouldn't be called if
	// it isn't a legal move
	b.position[col][board.colCount[col]] = piece
	b.colCount[col]++

	b.turn = p //Adjust the last turn to the current player
	if board.IsGameOver() {
		fmt.Println("THAT'S THE GAME FOLKS!")

		if board.IsDraw() {
			fmt.Println("IT'S A DRAW!")
		} else {
			fmt.Printf("THE WINNER IS: %s\n", p.Name)
		}
		fmt.Println("%s\nFinal Board Position:%s\n")
		fmt.Println(board.String())
	}

	b.adjustTurn(p)
	return b
}

// LegalMoves returns all of the current legal moves.
// Remember, a move is just the column you can play.
func (board C4Board) LegalMoves() []Move {
	// Creates a slice (start empty) to collect legal moves
	var legalMoves []Move

	// Appends a possible move if it isn't full
	var i uint
	for i = 0; i < board.numCols; i++ {
		if board.colCount[i] < board.numRows {
			legalMoves = append(legalMoves, Move(i))
		}
	}

	return legalMoves
}

// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// Evaluations on the board current state to determine winners and losers
// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// ------------------------------------------------------
// Function for telling if the game is currently over based upon the evaluations
// of the current Board
func (board C4Board) IsGameOver() bool {
	return board.IsWin() || board.IsDraw()
}

// IsWin calculates if the board is in a winning position
// if it is, then returns true, else returns false.
func (board C4Board) IsWin() bool {
	// Checks if there is a win in any direction
	if board.HorizontalWin() || board.VerticalWin() || board.DiagonalWin() {
		return true
	}
	return false
}

// IsDraw determines if the board is currently in a draw state
func (board C4Board) IsDraw() bool {

	// If there are no legal moves AND it isn't currently a win, then
	// its a draw. Theoretically, IsDraw is never called before IsWin, therefore
	// we know the board isn't in a winning state and don't neccesarily need that check.
	if len(board.LegalMoves()) == 0 && !board.IsWin() {
		return true
	}

	return false
}

// Evaluate returns the value of the piece's board
// This function scores the position for player
// and returns a numerical score
// When player is doing well, the score should be higher
// When player is doing worse, player's returned score should be lower
// Scores mean nothing except in relation to one another; so you can
// use any scale that makes sense to you
// The more accurately Evaluate() scores a position, the better that minimax will work
// There may be more than one way to evaluate a position but an obvious route
// is to count how many 1 filled, 2 filled, and 3 filled segments of the board
// that the player has (that don't include any of the opponents pieces) and give
// a higher score for 3 filleds than 2 filleds, 1 filleds, etc.
// You may also need to score wins (4 filleds) as very high scores and losses (4 filleds
// for the opponent) as very low scores
func (board C4Board) Evaluate(player Piece) float32 {
	var totalScore float32

	// These will load all of the segments for each direction into these three variables
	horizontalSegments, _ := board.CheckHorizontal()
	verticalSegments, _ := board.CheckVertical()
	diagonalSegments, _ := board.CheckDiagonal()

	// Gets the score for all the segments in that direction
	totalScore += CalculateDirection(horizontalSegments, player)
	totalScore += CalculateDirection(verticalSegments, player)
	totalScore += CalculateDirection(diagonalSegments, player)

	return totalScore
}

// segmentEquivalent checks if all of the pieces in the segment
// are of the same kind and non empty. Returns true if all pieces
// in the segment are of the same kind, false otherwise.
func segmentEquivalent(segment Segment) bool {
	// basePiece will be set to the first item in the slice,
	// doesn't matter if the first item is empty, because if it is
	// then all cannot be equivalent AND have a piece in it.
	basePiece := segment[0]
	for _, piece := range segment {
		if piece != basePiece || piece == Empty {
			return false
		}
	}

	return true
}

// CheckVertical checks if there is a winning vertical segment
// it will return immediately if a win is found. If a win is not found
// it will return all the segments tested and a win status of "false".
func (board C4Board) CheckVertical() (segments []Segment, win bool) {
	win = false
	var segment Segment

	// Finds all vertical segments and appends to a slice
	var i, j uint
	for i = 0; i < board.numCols; i++ {
		for j = 0; j < board.numRows-3; j++ {
			segment = Segment{
				board.position[i][j],
				board.position[i][j+1],
				board.position[i][j+2],
				board.position[i][j+3],
			}
			segments = append(segments, segment)

			if segmentEquivalent(segment) {
				win = true
			}
		}
	}

	return
}

// VerticalWin makes checking for winning conditions
// on the board look cleaner for vertical checking
func (board C4Board) VerticalWin() bool {
	_, verticalWin := board.CheckVertical()

	return verticalWin
}

// CheckHorizontal checks if there is a winning vertical segment
// it will return immediately if a win is found. If a win is not found
// it will return all the segments tested and a win status of "false".
func (board C4Board) CheckHorizontal() (segments []Segment, win bool) {
	win = false
	var segment Segment

	var i, j uint

	// Finds all horizontal segments and appends to a slice
	for i = 0; i < board.numRows; i++ {
		for j = 0; j < board.numCols-3; j++ {
			segment = Segment{
				board.position[j][i],
				board.position[j+1][i],
				board.position[j+2][i],
				board.position[j+3][i],
			}
			segments = append(segments, segment)

			if segmentEquivalent(segment) {
				win = true
			}
		}
	}

	return
}

// HorizontalWin makes checking for winning conditions
// on the board look cleaner for horizontal checking
func (board C4Board) HorizontalWin() bool {
	_, horizontalWin := board.CheckHorizontal()

	return horizontalWin
}

// CheckDiagonal checks if there is a winning diagonal segment
// it will return immediately if a win is found. If a win is not found
// it will return all the segments tested and a win status of "false".
func (board C4Board) CheckDiagonal() (segments []Segment, win bool) {
	win = false
	var segment Segment

	var i, j uint
	// Left to right diagonal checking
	for i = 0; i < board.numCols-3; i++ {
		for j = 0; j < board.numRows-3; j++ {
			segment = Segment{
				board.position[i][j],
				board.position[i+1][j+1],
				board.position[i+2][j+2],
				board.position[i+3][j+3],
			}
			segments = append(segments, segment)

			if segmentEquivalent(segment) {
				win = true
			}
		}
	}

	// Right to left diagonal checking
	for i = board.numCols - 1; i > 2; i-- {
		for j = 0; j < board.numRows-3; j++ {
			segment = Segment{
				board.position[i][j],
				board.position[i-1][j+1],
				board.position[i-2][j+2],
				board.position[i-3][j+3],
			}
			segments = append(segments, segment)

			if segmentEquivalent(segment) {
				win = true
			}
		}
	}

	return
}

// DiagonalWin makes checking for winning conditions
// on the board look cleaner for diagonal checking
func (board C4Board) DiagonalWin() bool {
	_, diagonalWin := board.CheckDiagonal()

	return diagonalWin
}

// CalculateDirection calculates the score in the direction of segments
func CalculateDirection(segments []Segment, player Piece) (score float32) {

	// Goes through every segment in the direction and
	// calculates the score for that segment
	for _, segment := range segments {
		score += CalculateScore(segment, player)
	}

	return
}

// Need to rewrite
func CalculateScore(segment Segment, player Piece) float32 {
	pieceCount := 0
	pieceToCount := Empty
	// Loops through all the pieces in the segment
	for _, piece := range segment {
		// We only want to choose a piece to count once
		// we actually get to a piece that isn't empty
		if piece != Empty && pieceToCount == Empty {
			pieceToCount = piece
			pieceCount++
		} else if piece != pieceToCount && piece != Empty {
			return 0.0
		} else if piece != Empty {
			pieceCount++
		}
	}

	// Closure to handle score calculating
	score := func() float32 {
		if pieceCount == 0 {
			return 0.0
		} else if pieceCount == 1 {
			return 1.0
		} else if pieceCount == 2 {
			return 5.0
		} else if pieceCount == 3 {
			return 50.0
		} else {
			return 5000.0
		}
	}
	if pieceToCount != player && pieceToCount != Empty {

		return -score()
	}

	return score()
}
