// Created by OrkWard at 2024/11/07 14:01
// leetgo: 1.4.10
// https://leetcode.com/problems/spiral-matrix-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func generateMatrix(n int) (ans [][]int) {
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	up, down, left, right := 0, n-1, 0, n-1

	count := 1
	for up <= right {
		// up line
		for i := left; i <= right; i += 1 {
			ans[up][i] = count
			count += 1
		}
		up += 1

		for i := up; i <= down; i += 1 {
			ans[i][right] = count
			count += 1
		}
		right -= 1

		if up <= down {
			for i := right; i >= left; i -= 1 {
				ans[down][i] = count
				count += 1
			}
			down -= 1
		}

		if left <= right {
			for i := down; i >= up; i -= 1 {
				ans[i][left] = count
				count += 1
			}
			left += 1
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := generateMatrix(n)

	fmt.Println("\noutput:", Serialize(ans))
}
