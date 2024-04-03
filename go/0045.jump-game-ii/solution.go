// Created by OrkWard at 2024/04/04 00:09
// leetgo: dev
// https://leetcode.com/problems/jump-game-ii/

package main

import (
  "bufio"
  "fmt"
  "os"

  . "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func jump(nums []int) (ans int) {
  jumpCount, curJumpTar, maxJumpTar := 0, 0, 0
  for i := 0; i < len(nums) - 1; i += 1 {
    if nums[i] + i > maxJumpTar {
      maxJumpTar = nums[i] + i
    }

    if curJumpTar == i {
      curJumpTar, jumpCount = maxJumpTar, jumpCount + 1
      maxJumpTar = nums[i] + i

      if curJumpTar >= len(nums) - 1 {
        return jumpCount
      }
    }
  }

  // never get here
  return 0
}

// @lc code=end

func main() {
  stdin := bufio.NewReader(os.Stdin)
  nums := Deserialize[[]int](ReadLine(stdin))
  ans := jump(nums)

  fmt.Println("\noutput:", Serialize(ans))
}
