// Created by OrkWard at 2024/11/15 09:33
// leetgo: 1.4.10
// https://leetcode.com/problems/plus-one/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func plusOne(digits []int) (ans []int) {
	carriage := 0
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i >= 0; i -= 1 {
		sum := digits[i] + carriage
		d := sum % 10
		carriage = sum / 10
		ans = append([]int{d}, ans...)
	}

	if carriage > 0 {
		ans = append([]int{1}, ans...)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	digits := Deserialize[[]int](ReadLine(stdin))
	ans := plusOne(digits)

	fmt.Println("\noutput:", Serialize(ans))
}
