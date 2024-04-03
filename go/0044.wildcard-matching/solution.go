// Created by OrkWard at 2024/04/04 00:28
// leetgo: dev
// https://leetcode.com/problems/wildcard-matching/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isMatch(s string, p string) bool {
    
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	p := Deserialize[string](ReadLine(stdin))
	ans := isMatch(s, p)

	fmt.Println("\noutput:", Serialize(ans))
}
