package logic

/*
Init the values array of the game board with '.' character.

return : pointer to the values game board array
*/
func InitValsArray() *[3][3]byte {
	return &[3][3]byte{{'.', '.', '.'}, {'.', '.', '.'}, {'.', '.', '.'}}
}

/*
Init the game board array with ' ' character.

return: pointer to the game board array
*/
func InitGameBoard() *[16][35]byte {
	boardArr := [16][35]byte{}
	for i := 0; i < 16; i++ {
		for j := 0; j < 35; j++ {
			if j == 11 || j == 12 || j == 22 || j == 23 {
				boardArr[i][j] = '*'
			} else if j >= 2 && j <= 32 && (i == 5 || i == 10) {
				boardArr[i][j] = '*'
			} else {
				boardArr[i][j] = ' '
			}
		}
	}
	return &boardArr
}

/*
Update the cells in the game board array regarding the given val.

parameters: val(byte) = the value that was inserted in the board(X/O)

			row(int) = the row which the value was inserted(0-2)

			col(int) = the col which the value was inserted(0-2)

			board(pointer to byte array) = pointer to array to the game board array
*/
func updateVals(val byte, row, col int, board *[16][35]byte) {
	switch val {
	case 'O':

		//update cells if O was inserted
		board[1+5*row][5+11*col] = '*'
		board[1+5*row][7+11*col] = '*'

		board[2+5*row][4+11*col] = '*'
		board[2+5*row][8+11*col] = '*'

		board[3+5*row][4+11*col] = '*'
		board[3+5*row][8+11*col] = '*'

		board[4+5*row][5+11*col] = '*'
		board[4+5*row][7+11*col] = '*'
	case 'X':

		//update cells if X was inserted
		board[1+5*row][3+12*col] = '*'
		board[1+5*row][4+12*col] = '*'
		board[1+5*row][6+12*col] = '*'
		board[1+5*row][7+12*col] = '*'

		board[2+5*row][4+12*col] = '*'
		board[2+5*row][5+12*col] = '*'
		board[2+5*row][6+12*col] = '*'

		board[3+5*row][4+12*col] = '*'
		board[3+5*row][5+12*col] = '*'
		board[3+5*row][6+12*col] = '*'

		board[4+5*row][3+12*col] = '*'
		board[4+5*row][4+12*col] = '*'
		board[4+5*row][6+12*col] = '*'
		board[4+5*row][7+12*col] = '*'
	}
}
