package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	former := countAndSay(n - 1)
	counter := 1
	currNum := former[0]
	say := ""

	for _, num := range former[1:] {
		if currNum == byte(num) {
			counter += 1
		} else {
			say += strconv.Itoa(counter)
			say += string(currNum)

			counter = 1
			currNum = byte(num)
		}
	}

	say += strconv.Itoa(counter)
	say += string(currNum)

	return say
}

// @lc code=end
