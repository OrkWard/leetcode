// Created by OrkWard at 2024/11/08 20:15
// leetgo: 1.4.10
// https://leetcode.com/problems/permutation-sequence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func getKthPermution(nums []int, k int, base []int) []int {
	if len(nums) == 1 {
		return append(base, nums...)
	}

	fact := factorial(len(nums) - 1)

	var nextNumIdx, nextK int
	if k%fact == 0 {
		nextNumIdx = k/fact - 1
		nextK = fact
	} else {
		nextNumIdx = k / fact
		nextK = k % fact
	}
	nextBaseNum := nums[nextNumIdx]
	nextNums := append(nums[:nextNumIdx], nums[nextNumIdx+1:]...)

	return getKthPermution(nextNums, nextK, append(base, nextBaseNum))
}

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := range n {
		nums[i] = i + 1
	}

	permutation := getKthPermution(nums, k, []int{})

	permutationStr := ""
	for _, num := range permutation {
		permutationStr += string(num + '0')
	}

	return permutationStr
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := getPermutation(n, k)

	fmt.Println("\noutput:", Serialize(ans))
}
