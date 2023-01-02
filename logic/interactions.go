package logic

import "fmt"

/*
Get the player name from the user.

return: first (string) = the player first name

*/
func GetPlayerName() string {
	var first string
	fmt.Scanln(&first)
	return first
}

/*
Get the chosen side from the player.

If invalid then returns an error, else a boolean var(true if the side is X)
*/
func GetSide() (bool, error) {
	var side string
	fmt.Scanln(&side)
	playerSide, err := checkPlayerSide(side)
	if err != nil {
		return false, fmt.Errorf("Unavailable Side, please choose O or X or o or x")
	} else {
		if playerSide == "O" {
			return false, nil
		}
		return true, nil
	}
}

/*
A player move function, first it reads the selected cell from the player and validate it.

if it invalid move than it will return an error, if it valid it will update the matched element

of the values Array and the matched elements in the game board array
*/
func PlayerMove(valArr *[3][3]byte, side bool, board *[16][35]byte) error {
	var cell byte
	var r, c int
	var err error
	//get cell from player
	fmt.Scanln(&cell)

	//validate cell
	r, c, err = validateMove(cell, valArr)

	//update values array and game board array
	if err == nil {

		var letter byte
		if side {
			letter = 'X'
		} else {
			letter = 'O'
		}
		valArr[r][c] = letter
		updateVals(letter, r, c, board)
	}
	return err
}
