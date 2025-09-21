// Main game function for looping throught the game state
package connect4

import (
	"fmt"

	board "github.com/accal/GoLangProjects/Connect4"
	player "github.com/accal/GoLangProjects/Connect4"
)

// Main function to play the Connect 4 game from list of programs
// This function is designed to let you play against a single CPU that will predict the best moves possible against you
// this is done by a MiniMax algorithm that will look ahead a certain depth to determine the best move
// This is a refactor of previous work done in the past for an assignment to create a simple Connect4 Game
func PlayConnect4() {
	fmt.Println("------------- Initializing Connect 4 -------------")
	displayDirections()
	//Define the initial game board setting the current player's turn to Black
	var gameBoard board.Board = board.C4Board{turn: player.PlayerIcon}

	//Assign the Players
	var player1 player.Player = player.Player{Name: "Player", TurnCount: player.incrementer(), Piece: player.PlayerIcon, IsHuman: true}
	var player2 player.Player = player.Player{Name: "Computer", TurnCount: player.incrementer(), Piece: player.CpuIcon, IsHuman: false}

	//Main Loop for the game until there is a win or a draw
	for !gameBoard.IsGameOver() {
		fmt.Println("%s\nCurrent Board:%s\n")
		fmt.Printf("%s", gameBoard.String())

		gameBoard = gameBoard.MakeMove(player1, col)
		player1.TurnCount()
		if gameBoard.IsGameOver() {
			break
		}

		gameBoard = gameBoard.MakeMove(player2, ConcurrentFindBestMove(gameBoard, 3)) //Concurrent without inputted Depth
		player2.TurnCount()
		if gameBoard.IsGameOver() {
			break
		}
	}
}

func displayDirections() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("---------------- Game Directions -----------------")
	fmt.Println("--------------------------------------------------")
	fmt.Println("You are playing as Black (X) and the Computer is Red (O)")
	fmt.Println("To make a move, enter the column number (0-6) where you want to drop your piece.")
	fmt.Println("--------------------------------------------------")

}
