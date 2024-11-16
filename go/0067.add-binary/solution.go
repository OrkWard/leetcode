// Created by OrkWard at 2024/11/17 00:52
// leetgo: dev
// https://leetcode.com/problems/add-binary/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func addBinary(a string, b string) (ans string) {
	i, j := len(a)-1, len(b)-1
	carry := 0
	for i >= 0 || j >= 0 || carry > 0 {
		sum :=  carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		carry = sum / 2
		ans = string(sum%2 + '0') + ans
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	a := Deserialize[string](ReadLine(stdin))
	b := Deserialize[string](ReadLine(stdin))
	ans := addBinary(a, b)

	fmt.Println("\noutput:", Serialize(ans))
}
