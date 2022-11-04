package logic

import (
	"math"
)

/*
The Ai object, fields:
board = pointer to 3x3 array that represent the game state after the player turn
playerSide = the side of the player (X/O)
compSide = the side of the computer(opposite of playerSide)
*/
type MiniMax struct {
	board      *[3][3]byte
	playerSide byte
	compSide   byte
}

/*
The indexes of a valid move in each possible scenario
*/
type move struct {
	r int
	c int
}

/*
Init the AI fields after each player turn
*/
func (AI *MiniMax) construct(isX bool, board *[3][3]byte) {
	if isX {
		AI.playerSide = 'X'
		AI.compSide = 'O'
	} else {
		AI.playerSide = 'O'
		AI.compSide = 'X'
	}
	AI.board = board
}

/*
Find the optimal move for the computer after each player move,

and update the relevant cells in the values game array and the game board array.

Return the outcome of the move(-10 computer wins, 10 player wins, 0 tie)
*/
func (AI *MiniMax) AIMove(isX bool, round int, board *[3][3]byte, screen *[16][35]byte) int {
	AI.construct(isX, board)
	r, c := AI.findBestMove(AI.board, round)
	board[r][c] = AI.compSide
	updateVals(AI.compSide, r, c, screen)
	return (AI).Evaluation(false, board)
}

/*
Checks if there is a winning situation(i.e an entire row or col or diagonal with the same sign)

If there is than it will return either 10 for player winning or -10 for the computer.

If there still is a tie then it will return 0
*/
func (AI *MiniMax) Evaluation(isPlayer bool, valArr *[3][3]byte) int {
	// row check
	for i := 0; i < 3; i++ {
		if valArr[i][0] == valArr[i][1] && valArr[i][0] == valArr[i][2] && valArr[i][0] != '.' {
			if isPlayer {
				return 10
			}
			return -10
		}
	}
	// col check
	for j := 0; j < 3; j++ {
		if valArr[0][j] == valArr[1][j] && valArr[0][j] == valArr[2][j] && valArr[0][j] != '.' {
			if isPlayer {
				return 10
			}
			return -10
		}
	}
	// diagonals checks
	if ((valArr[0][0] == valArr[2][2] && valArr[0][0] == valArr[1][1]) ||
		(valArr[0][2] == valArr[2][0] && valArr[0][2] == valArr[1][1])) && valArr[1][1] != '.' {
		if isPlayer {
			return 10
		}
		return -10
	}
	return 0
}

/*
Return the indexes of the next optimal move for the computer
*/
func (AI *MiniMax) findBestMove(board *[3][3]byte, round int) (r int, c int) {

	//game state to minimize
	toMinimize := math.MaxInt
	//find the possible moves
	moves := getMoves(board)
	//find the optimal move
	for _, m := range moves {
		//play move
		board[m.r][m.c] = AI.compSide
		//find score of the move
		val := AI.minimax(AI.board, 1, true, false, round+1, math.MinInt, math.MaxInt)
		//undue move
		board[m.r][m.c] = '.'
		//if optimal so far
		if toMinimize > val {
			r = m.r
			c = m.c
			toMinimize = val
		}
	}
	return r, c
}

/*
Simulate the outcomes of a given move and return the score of the final game state of the move
*/
func (AI *MiniMax) minimax(board *[3][3]byte, depth int, isMaximizingPlayer bool, isPlayer bool, round int, alpha, beta int) int {

	//evaluate the game state
	score := AI.Evaluation(isPlayer, board)
	//if player won
	if score > 0 {
		return score - depth
		//if computer won
	} else if score < 0 {
		return score + depth
		//if there is a tie and also no more moves are available
	} else if AI.isMovesLeft(round) == false {
		return 0
		//there are still more scenario's to check
	} else {
		//if player turn
		if isMaximizingPlayer {
			//game score to maximize
			bestVal := math.MinInt
			//find possible player moves
			moves := getMoves(board)
			for _, m := range moves {
				//play move
				board[m.r][m.c] = AI.playerSide
				//find the final score of this move
				value := AI.minimax(board, depth+1, false, true, round+1, alpha, beta)
				//undue move
				board[m.r][m.c] = '.'
				//if optimal so far
				if value > bestVal {
					bestVal = value
				}
				//maximize alpha
				if alpha < bestVal {
					alpha = bestVal
				}
				//pruning
				if beta <= alpha {
					break
				}
			}
			//return optimal score
			return bestVal
		} else {
			//game score to minimize
			bestVal := math.MaxInt
			//find possible computer moves
			moves := getMoves(board)
			for _, m := range moves {
				//play move
				board[m.r][m.c] = AI.compSide
				//find the final score of this move
				value := AI.minimax(board, depth+1, true, false, round+1, alpha, beta)
				//undue move
				board[m.r][m.c] = '.'
				//if optimal so far
				if value < bestVal {
					bestVal = value
				}
				//minimize beta
				if beta > bestVal {
					beta = bestVal
				}
				//pruning
				if beta <= alpha {
					break
				}
			}
			//return optimal score
			return bestVal
		}
	}
}

/*
Return a slice of moves objects of possible moves given a state of the game board
*/
func getMoves(board *[3][3]byte) []move {
	list := make([]move, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == '.' {
				var m move
				m.r = i
				m.c = j
				list = append(list, m)
			}
		}
	}
	return list
}

/*
Determine if there are still moves to play
*/
func (AI *MiniMax) isMovesLeft(round int) bool {
	if round < 9 {
		return true
	}
	return false
}
