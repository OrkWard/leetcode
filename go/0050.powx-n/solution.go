// Created by OrkWard at 2024/06/14 00:51
// leetgo: dev
// https://leetcode.com/problems/powx-n/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func myPow(x float64, n int) (ans float64) {
	if x == 0 {
		return 0
	} else if n > 0 {
		ans = 1.0
		for _ = range n {
			ans *= x
		}
		return ans
	} else {
		ans = 1.0
		for _ = range -n {
			ans /= x
			if (ans == 0.0) {
				return 0.0
			}
		}
		return ans
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[float64](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := myPow(x, n)

	fmt.Println("\noutput:", Serialize(ans))
}
