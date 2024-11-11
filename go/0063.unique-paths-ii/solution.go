// Created by OrkWard at 2024/11/11 10:11
// leetgo: 1.4.10
// https://leetcode.com/problems/unique-paths-ii/

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

func uniquePathsWithMN(obstacleGrid [][]int, m int, n int) (ans int) {
	if m == 1 {
		for i := range n {
			if obstacleGrid[0][i] == 1 {
				return 0
			}
		}
		return 1
	}

	if n == 1 {
		for i := range m {
			if obstacleGrid[i][0] == 1 {
				return 0
			}
		}
		return 1
	}

	key := strconv.Itoa(m-1) + "," + strconv.Itoa(n-1)
	v, e := cache[key]
	if e {
		return v
	}

	if obstacleGrid[m-2][n-1] == 0 {
		ans += uniquePathsWithMN(obstacleGrid, m-1, n)
	}
	if obstacleGrid[m-1][n-2] == 0 {
		ans += uniquePathsWithMN(obstacleGrid, m, n-1)
	}

	cache[key] = ans
	return
}

func uniquePathsWithObstacles(obstacleGrid [][]int) (ans int) {
	cache = make(map[string]int)
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	return uniquePathsWithMN(obstacleGrid, m, n)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	obstacleGrid := Deserialize[[][]int](ReadLine(stdin))
	ans := uniquePathsWithObstacles(obstacleGrid)

	fmt.Println("\noutput:", Serialize(ans))
}
