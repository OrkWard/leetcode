// Created by OrkWard at 2024/11/13 15:41
// leetgo: 1.4.10
// https://leetcode.com/problems/valid-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
import (
	"strings"
	"unicode"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}

	seenExp := false
	seenDot := false
	seenDigit := false

	for i, r := range s {
		if unicode.IsDigit(r) {
			seenDigit = true
		} else if r == '.' {
			if seenDot || seenExp {
				return false
			}

			seenDot = true
		} else if r == 'e' || r == 'E' {
			if !seenDigit || seenExp {
				return false
			}

			seenExp = true
			seenDigit = false
		} else if r == '-' || r == '+' {
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			return false
		}
	}

	if !seenDigit {
		return false
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := isNumber(s)

	fmt.Println("\noutput:", Serialize(ans))
}
