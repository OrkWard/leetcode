// Created by OrkWard at 2024/11/06 13:59
// leetgo: 1.4.10
// https://leetcode.com/problems/length-of-last-word/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLastWord(s string) (ans int) {
	end := len(s)-1
	for i := len(s)-1; i >= 0; i -= 1 {
		if s[i] != ' ' {
			end = i // must > 0, for at least one word
			break
		}
	}

	for i := end; i >= 0; i -= 1 {
		if s[i] == ' ' {
			return end - i
		}
	}

	return end+1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLastWord(s)

	fmt.Println("\noutput:", Serialize(ans))
}
