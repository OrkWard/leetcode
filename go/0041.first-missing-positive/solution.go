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

func abs(num int) int {
  if num < 0 {
    return -num
  }
  return num
}

func firstMissingPositive(nums []int) (ans int) {
  for i := range len(nums) {
    if nums[i] <= 0 {
      nums[i] = len(nums) + 1
    }
  }

  for i := range len(nums) {
    if nums[i] >= -len(nums) && nums[i] <= len(nums) {
      // set flag
      numExist := abs(nums[i])
      nums[numExist - 1] = -abs(nums[numExist -1])
    }
  }

  for i := range len(nums) {
    if nums[i] > 0 {
      return i + 1
    }
  }
  return len(nums) + 1
}

// @lc code=end

func main() {
  stdin := bufio.NewReader(os.Stdin)
  nums := Deserialize[[]int](ReadLine(stdin))
  ans := firstMissingPositive(nums)

  fmt.Println("\noutput:", Serialize(ans))
}
