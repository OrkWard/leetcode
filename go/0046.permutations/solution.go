// Created by OrkWard at 2024/04/06 19:57
// leetgo: dev
// https://leetcode.com/problems/permutations/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func permute(nums []int) (ans [][]int) {
  var res [][]int
  backtrace([]int{}, nums, &res)
	return res
}

func backtrace(path, choices []int, res *[][]int) {
  if 0 == len(choices) {
    *res = append(*res, path)
    return
  }

  for i, l := range choices {
    backtrace(
      append(path, l),
      append(append([]int{}, choices[:i]...), choices[i+1:]...),
      res,
    )
  }
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := permute(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
