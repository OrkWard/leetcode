package main

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	var numbers = [10]byte{}

	// format
	if len(board) != 9 {
		return false
	}

	for i := 0; i < 9; i += 1 {
		if len(board[i]) != 9 {
			return false
		}
		for j := 0; j < 9; j += 1 {

			if board[i][j] < '0' {
				board[i][j] = 0
			} else {

				board[i][j] -= '0'
			}
		}
	}

	// fmt.Println("row")

	// validate rows
	for _, row := range board {
		for _, item := range row {
			if item > 9 {
				return false
			}

			numbers[item] += 1
		}

		// fmt.Println(numbers)
		if !validateSeq(&numbers) {
			return false
		}
	}

	// fmt.Println("columns")
	// validate columns
	for i := 0; i < 9; i += 1 {
		for j := 0; j < 9; j += 1 {
			numbers[board[j][i]] += 1
		}

		// fmt.Println(numbers)
		if !validateSeq(&numbers) {
			return false
		}

	}

	// validate boxes
	coords := [9][2]int8{{1, 1}, {1, 0}, {1, -1},
		{0, 1}, {0, 0}, {0, -1},
		{-1, 1}, {-1, 0}, {-1, -1}}

	for _, box := range coords {
		center := [2]int8{4, 4}
		center[0] += box[0] * 3
		center[1] += box[1] * 3
		for _, offset := range coords {
			numbers[board[center[0]+offset[0]][center[1]+offset[1]]] += 1
		}

		// fmt.Println(numbers)

		if !validateSeq(&numbers) {
			return false
		}

	}

	return true
}

func validateSeq(seq *[10]byte) bool {
	for i, item := range seq[1:10] {
		if item > 1 {
			return false
		} else {
			seq[i+1] = 0
		}
	}

	return true
}

// @lc code=end
