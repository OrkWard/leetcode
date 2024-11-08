// Created by OrkWard at 2024/10/22 10:53
// leetgo: 1.4.9
// https://leetcode.com/problems/spiral-matrix/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func spiralOrder(matrix [][]int) (ans []int) {
	if (len(matrix) == 0) {
		return make([]int, 0)
	}

	row, col := len(matrix), len(matrix[0])
	top, right, bottom, left := 0, col-1, row-1, 0

	result := make([]int, 0, col * row)

	for top <= bottom && left <= right {
		for i := left ; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}
	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	ans := spiralOrder(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}
