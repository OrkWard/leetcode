// Created by OrkWard at 2024/10/23 11:08
// leetgo: 1.4.9
// https://leetcode.com/problems/jump-game/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canJump(nums []int) bool {
	candidates := make([]int, len(nums))
	candidates[len(nums)-1] = 1
	for i := len(nums)-1; i >= 0; i -= 1 {
		if candidates[i] != 1 {
			continue
		}

		for j := 0; j < i; j += 1 {
			if j + nums[j] >= i {
				candidates[j] = 1
			}
		}
	}

	return candidates[0] == 1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := canJump(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
