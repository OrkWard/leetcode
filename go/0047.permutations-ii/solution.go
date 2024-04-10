// Created by OrkWard at 2024/04/07 17:37
// leetgo: dev
// https://leetcode.com/problems/permutations-ii/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func permuteUnique(nums []int) (ans [][]int) {
	var permutations [][]int
	sort.Ints(nums)
	permute([]int{}, nums, &permutations)
	return permutations
}

func permute(path []int, choices []int, permutations *[][]int) {
	if len(choices) == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*permutations = append(*permutations, temp)
		return
	}

	for i, num := range choices {
		if i == 0 || num != choices[i-1] {
			permute(
				append(path, num),
				append(append([]int{}, choices[:i]...), choices[i+1:]...), permutations)
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := permuteUnique(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
