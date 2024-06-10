// Created by OrkWard at 2024/06/11 00:03
// leetgo: dev
// https://leetcode.com/problems/group-anagrams/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type Key [26]byte

func strtoKey(str string) (key Key) {
	for i := range str {
		key[str[i] - 'a']++
	}

	return
}

func groupAnagrams(strs []string) (ans [][]string) {
	groups := make(map[Key][]string)

	for _, str := range strs {
		key := strtoKey(str)
		groups[key] = append(groups[key], str)
	}

	ans = make([][]string, 0, len(groups))

	for _, grp := range groups {
		ans = append(ans, grp)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	strs := Deserialize[[]string](ReadLine(stdin))
	ans := groupAnagrams(strs)

	fmt.Println("\noutput:", Serialize(ans))
}
