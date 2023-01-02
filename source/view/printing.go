package view

import "fmt"

/*
Prints the game board after each round.

parameters: board(pointer to byte array) = pointer to array to the game board array.
*/
func PrintBoard(board *[16][35]byte) {
	for i := 0; i < 16; i++ {
		for j := 0; j < 35; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

//printing messages to player functions

func PrintingStartMessage(name string) {
	fmt.Printf("Hello %s, welcome to the game.\n", name)
	fmt.Printf("Please choose side(O, X): ")
}

func PrintRequestNameMessage() {
	fmt.Printf("Please enter your name: ")
}

func PrintMessage(msg string) {
	fmt.Println(msg)
}

func PrintSideMessage(side bool) {
	var s string
	if side {
		s = "X"
	} else {
		s = "O"
	}
	fmt.Printf("You are %s\n", s)
}

func PrintPlayerTurnMsg(name string) {
	fmt.Printf("%s turn, please select cell(1-9): ", name)
}

func PrintComputerTurnMsg() {
	fmt.Println("Computer turn ....")
}
