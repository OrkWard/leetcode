// Created by OrkWard at 2024/08/09 10:25
// leetgo: 1.4.7
// https://leetcode.com/problems/maximum-subarray/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxSubArray(nums []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxSubArray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
