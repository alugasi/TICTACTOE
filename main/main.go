package main

import (
	l "TICTACTOE/logic"
	v "TICTACTOE/view"
)

func main() {

	winner := 0 //0 = no winner, negative = computer wins, positive = player wins
	vArr := l.InitValsArray()
	//init game board
	bArr := l.InitGameBoard()
	//ask for player name
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

	//create AI object
	var computer l.MiniMax

	//start game
	for i := 0; i < 9 && winner == 0; {
		//player turn
		v.PrintPlayerTurnMsg(playerName)
		//making player move
		row, col, playerMoveErr := l.PlayerMove(vArr, isX, bArr)
		for playerMoveErr != nil {
			v.PrintMessage(playerMoveErr.Error())
			v.PrintPlayerTurnMsg(playerName)
			row, col, playerMoveErr = l.PlayerMove(vArr, isX, bArr)
		}
		v.PrintBoard(bArr)
		i++
		//validate player turn
		winner = (&computer).Evaluation(row, col, isX, 1, vArr)

		if winner == 0 && i < 9 {
			//computer turn
			v.PrintComputerTurnMsg()
			//making computer move
			(&computer).AIMove(isX, i, vArr, bArr)
			//validate computer turn
			winner = (&computer).Evaluation(row, col, !isX, 0, vArr)
			v.PrintBoard(bArr)
			i++
		}
	}
	//print outcome
	if winner == 0 {
		v.PrintMessage("Tie!")
	} else if winner > 0 {
		v.PrintMessage("Player wins!")
	} else {
		v.PrintMessage("Computer wins!")
	}
}
