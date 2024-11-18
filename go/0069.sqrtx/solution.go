// Created by OrkWard at 2024/11/18 10:43
// leetgo: 1.4.10
// https://leetcode.com/problems/sqrtx/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mySqrt(x int) (ans int) {
	if x == 0 {
		return 0
	}
	if x <= 3 {
		return 1
	}

	l, r := 0, x
	for r - l > 1 {
		m := (l+r) / 2
		if m*m == x {
			return m
		} else if m*m > x {
			r = m
		} else {
			l = m
		}
	}
	return l
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := mySqrt(x)

	fmt.Println("\noutput:", Serialize(ans))
}
