// Created by OrkWard at 2024/11/18 19:20
// leetgo: 1.4.10
// https://leetcode.com/problems/climbing-stairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
var cache map[int]int

func climbStairs(n int) (ans int) {
	if n <= 2 {
		return n
	}

	if cache == nil {
		cache = make(map[int]int)
	}

	if v, exist := cache[n]; exist {
		return v
	}

	ans = climbStairs(n-1) + climbStairs(n-2)
	cache[n] = ans

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := climbStairs(n)

	fmt.Println("\noutput:", Serialize(ans))
}
