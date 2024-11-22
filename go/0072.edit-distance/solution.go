// Created by OrkWard at 2024/11/22 10:37
// leetgo: 1.4.10
// https://leetcode.com/problems/edit-distance/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minDistance(word1 string, word2 string) (ans int) {
	m, n := len(word1), len(word2)
	cache := make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i += 1 {
		cache[i][0] = i
	}
	for i := 0; i <= n; i += 1 {
		cache[0][i] = i
	}

	for i := 1; i <= m; i += 1 {
		for j := 1; j <= n; j += 1 {
			if word1[i-1] == word2[j-1] {
				cache[i][j] = cache[i-1][j-1]
			} else {
				cache[i][j] = min(cache[i-1][j], cache[i][j-1], cache[i-1][j-1]) + 1
			}
		}
	}

	return cache[m][n]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	word1 := Deserialize[string](ReadLine(stdin))
	word2 := Deserialize[string](ReadLine(stdin))
	ans := minDistance(word1, word2)

	fmt.Println("\noutput:", Serialize(ans))
}
