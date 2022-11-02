package main

import (
	l "TICTACTOE/logic"
	v "TICTACTOE/view"
)

func main() {

	winner := -1 //-1 = no winner, 0 = computer wins, 1 = player wins
	vArr := l.InitValsArray()
	//init game board
	bArr := l.InitGameBoard()
	//asl for player name
	v.PrintRequestNameMessage()
	playerName := l.GetPlayerName()
	var isX bool
	var getSideError error
	//get side from player
	v.PrintingStartMessage(playerName)
	isX, getSideError = l.GetSide()
	for getSideError != nil {
		v.PrintMessage(getSideError.Error())
		v.PrintingStartMessage(playerName)
		isX, getSideError = l.GetSide()
	}
	//announce player side
	v.PrintSideMessage(isX)
	//start game
	for i := 0; i < 9 && winner == -1; i++ {
		//player turn
		v.PrintPlayerTurnMsg(playerName)
		row, col, playerMoveErr := l.PlayerMove(vArr, isX, bArr)
		for playerMoveErr != nil {
			v.PrintMessage(playerMoveErr.Error())
			v.PrintPlayerTurnMsg(playerName)
			row, col, playerMoveErr = l.PlayerMove(vArr, isX, bArr)
		}
		v.PrintBoard(bArr)
		//validate player turn
		winner = l.CheckWinning(row, col, isX, 1, vArr)
		if winner == -1 {
			//computer turn
			v.PrintComputerTurnMsg()
			//validate computer turn
			winner = l.CheckWinning(row, col, isX, 0, vArr)
			v.PrintBoard(bArr)
		}
	}
	//print outcome
	if winner == 0 {
		v.PrintMessage("Tie!")
	} else if winner == 1 {
		v.PrintMessage("Player wins!")
	} else {
		v.PrintMessage("Computer wins!")
	}
}
