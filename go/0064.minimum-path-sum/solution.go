// Created by OrkWard at 2024/11/12 10:27
// leetgo: 1.4.10
// https://leetcode.com/problems/minimum-path-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
import "strconv"

var cache map[string]int

func minPathSumMN(grid [][]int, m int, n int) (sum int) {
	if m == 1 {
		for i := range n {
			sum += grid[0][i]
		}
		return
	}

	if n == 1 {
		for i := range m {
			sum += grid[i][0]
		}
		return
	}

	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	if v, e := cache[key]; e {
		return v
	}

	sum = min(minPathSumMN(grid, m-1, n), minPathSumMN(grid, m, n-1)) + grid[m-1][n-1]
	cache[key] = sum

	return
}

func minPathSum(grid [][]int) (ans int) {
	cache = make(map[string]int)
	m := len(grid)
	n := len(grid[0])

	return minPathSumMN(grid, m, n)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	ans := minPathSum(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
