// Created by OrkWard at 2024/08/08 19:24
// leetgo: 1.4.7
// https://leetcode.com/problems/n-queens-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type queen struct {
	row int
	col int
}

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}

/* only verify lastest */
func isValidQueens(queens []queen) bool {
	if len(queens) <= 1 {
		return true
	}

	lastQueen := queens[len(queens)-1]
	for _, q := range queens[:len(queens)-1] {
		if q.col == lastQueen.col || q.row == lastQueen.row || abs(q.col-lastQueen.col) == abs(q.row-lastQueen.row) {
			return false
		}
	}

	return true
}

func totalNQueens(n int) (ans int) {
	// fmt.Println("0> ")
	if n == 1 {
		ans = 1
		return ans
	}

	var queens []queen
	for i := range n {
		queens = []queen{{i, 0}}
		var backtrace func(col int)

		backtrace = func(col int) {
			// fmt.Println("3> ")
			if col == n-1 {
				for i := range n {
					queens = append(queens[:col], queen{i, col})
					// fmt.Println(queens)
					if isValidQueens(queens) {
						ans += 1
						return
					}
				}
				return
			}

		label:
			for i := range n {
				// fmt.Println("1> ", queen{i, col})
				queens = append(queens[:col], queen{i, col})
				if !isValidQueens(queens) {
					continue label
				}
				backtrace(col + 1)
			}
		}

		backtrace(1)
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := totalNQueens(n)

	fmt.Println("\noutput:", Serialize(ans))
}
