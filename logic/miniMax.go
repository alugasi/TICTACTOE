package logic

type MiniMax struct {
	board      [3][3]byte
	playerSide byte
	compSide   byte
	round      int
}

func (AI *MiniMax) construct(isX bool, round int, board *[3][3]byte) {
	if isX {
		AI.playerSide = 'X'
		AI.compSide = 'O'
	} else {
		AI.playerSide = 'O'
		AI.compSide = 'X'
	}
	AI.round = round
	AI.board = *board
}

func (AI *MiniMax) AIMove(isX bool, round int, board *[3][3]byte, screen *[16][35]byte) {
	AI.construct(isX, round, board)
	r, c := AI.minimax()
	board[r][c] = AI.compSide
	updateVals(AI.compSide, r, c, screen)
}

/*
Checks if there is a winning situation(i.e an entire row or col or diagonal with the same sign)

If there is than it will return either 10 for player winning or -10 for the computer.

If there still is a tie then it will return 0
*/
func (AI *MiniMax) Evaluation(row, col int, isX bool, turn int, valArr *[3][3]byte) int {
	var signToCheck byte
	if isX {
		signToCheck = 'X'
	} else {
		signToCheck = 'O'
	}
	//row and col check
	if (valArr[row][0] == signToCheck && valArr[row][1] == signToCheck && valArr[row][2] == signToCheck) ||
		(valArr[0][col] == signToCheck && valArr[1][col] == signToCheck && valArr[2][col] == signToCheck) {
		if turn == 1 {
			return 10
		}
		return -10
	}
	// diagonals checks
	if valArr[1][1] == signToCheck && ((valArr[0][0] == signToCheck && valArr[2][2] == signToCheck) ||
		(valArr[0][2] == signToCheck && valArr[2][0] == signToCheck)) {
		if turn == 1 {
			return 10
		}
		return -10
	}
	return 0
}

func (AI *MiniMax) findBestMove() {

}

func (AI *MiniMax) minimax() (r int, c int) {

	return r, c
}

func (AI *MiniMax) isMovesLeft(round int) bool {
	if round < 9 {
		return true
	}
	return false
}
