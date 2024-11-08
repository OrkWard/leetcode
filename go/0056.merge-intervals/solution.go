// Created by OrkWard at 2024/11/05 10:48
// leetgo: 1.4.10
// https://leetcode.com/problems/merge-intervals/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

import (
	"sort"
)

func merge(intervals [][]int) (ans [][]int) {
	if len(intervals) == 0 {
		return
	}

	sort.Slice(intervals, func (i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, intervals[0])
	for _, interval := range intervals[1:] {
		last := ans[len(ans)-1]
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			ans = append(ans, interval)
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	intervals := Deserialize[[][]int](ReadLine(stdin))
	ans := merge(intervals)

	fmt.Println("\noutput:", Serialize(ans))
}
