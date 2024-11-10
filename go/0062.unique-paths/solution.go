// Created by OrkWard at 2024/11/10 14:43
// leetgo: dev
// https://leetcode.com/problems/unique-paths/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
var cache = make(map[string]int)

func uniquePaths(m int, n int) (ans int) {
	if (m == 1 || n == 1) {
		return 1
	}

	v, e := cache[string(m) + "," + string(n)]
	if e {
		return v
	}
	ans = uniquePaths(m - 1, n) + uniquePaths(m, n - 1)
	cache[string(m) + "," + string(n)] = ans
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	m := Deserialize[int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := uniquePaths(m, n)

	fmt.Println("\noutput:", Serialize(ans))
}
