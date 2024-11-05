// Created by OrkWard at 2024/11/05 13:57
// leetgo: 1.4.10
// https://leetcode.com/problems/insert-interval/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	for i, interval := range intervals {
		if interval[1] < newInterval[0] {
			ans = append(ans, interval)
		} else if interval[0] > newInterval[1] {
			ans = append(ans, newInterval)
			ans = append(ans, intervals[i:]...)
			return
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}

	ans = append(ans, newInterval)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	intervals := Deserialize[[][]int](ReadLine(stdin))
	newInterval := Deserialize[[]int](ReadLine(stdin))
	ans := insert(intervals, newInterval)

	fmt.Println("\noutput:", Serialize(ans))
}
