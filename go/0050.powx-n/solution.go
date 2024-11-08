// Created by OrkWard at 2024/08/06 14:40
// leetgo: 1.4.7
// https://leetcode.com/problems/powx-n/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func myPow(x float64, n int) (ans float64) {
	if n == 0 {
		return 1
	}

	sign := n / abs(n)
	n = abs(n)
	ans = 1
	for n > 0 {
		mod := n % 2
		if mod == 1 && sign > 0 {
			ans *= x
		} else if mod == 1 && sign < 0 {
			ans /= x
		}
		x = x * x
		n = n / 2
	}
	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[float64](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := myPow(x, n)

	fmt.Println("\noutput:", Serialize(ans))
}
