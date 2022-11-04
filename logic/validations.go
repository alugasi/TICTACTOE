package logic

import "fmt"

/*
Validate if the side the player chose is valid (O/X/o/x).

If not it will return an error

Parameters: side(string) = The chosen side can be only one of the following: X/O/x/o.

return: side(string) = the chosen side ("" if not valid, O/X if valid).

		err = nil if valid else an error.
*/
func checkPlayerSide(side string) (string, error) {
	if side != "O" && side != "X" && side != "o" && side != "x" {
		return "", fmt.Errorf("Error")
	} else {
		if side == "o" {
			side = "O"
		} else if side == "x" {
			side = "X"
		}
		return side, nil
	}
}

/*
Validate a move of the player.

If it valid(cell is between 1-9 and not taken), it will return the indexes and nil error,

else return -1,-1 for the indexes and an error
*/
func validateMove(cell byte, valArr *[3][3]byte) (int, int, error) {
	//valid cell digit
	if cell >= 1 && cell <= 9 {
		index := int(cell)
		index = index - 1
		r, c := getIndexes(index)
		//cell is taken
		if valArr[r][c] != '.' {
			return -1, -1, fmt.Errorf("Cell is already taken!")
		} else {
			return r, c, nil
		}
	} else {
		return -1, -1, fmt.Errorf("Invalid cell(must be in range 1-9)!")
	}
}

/*

Return the array indexes (row and coll), based on the player chosen cell(0-8)

*/
func getIndexes(index int) (int, int) {
	switch index {
	case 0:
		return 0, 0
	case 1:
		return 0, 1
	case 2:
		return 0, 2
	case 3:
		return 1, 0
	case 4:
		return 1, 1
	case 5:
		return 1, 2
	case 6:
		return 2, 0
	case 7:
		return 2, 1
	default:
		return 2, 2
	}
}
