// Created by OrkWard at 2024/03/21 19:16
// leetgo: dev
// https://leetcode.com/problems/first-missing-positive/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func firstMissingPositive(nums []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := firstMissingPositive(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
