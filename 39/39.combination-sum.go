package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	track := []int{}

	backtrack(candidates, target, track, &res)

	return res
}

func backtrack(candidates []int, target int, track []int, res *[][]int) {
	sum := 0
	for _, num := range track {
		sum += num
	}

	if sum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		fmt.Println(temp, sum)
		*res = append(*res, temp)
		return
	} else if sum > target {
		return
	}

	for _, candidate := range candidates {
		if len(track) == 0 || candidate >= track[len(track)-1] {
			track = append(track, candidate)
			backtrack(candidates, target, track, res)
			track = track[:len(track)-1]
		}
	}
}

// @lc code=end
