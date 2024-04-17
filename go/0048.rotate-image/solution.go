// Created by OrkWard at 2024/04/10 16:58
// leetgo: dev
// https://leetcode.com/problems/rotate-image/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type pos struct {
	x int
	y int
}

func rotateSingle(point pos, matrix *[][]int) {
	pointF1 := pos{len(*matrix) - 1 - point.y, point.x}
	pointF2 := pos{len(*matrix) - 1 - point.x, len(*matrix) - 1 - point.y}
	pointF3 := pos{point.y, len(*matrix) - 1 - point.x}

	tmp := (*matrix)[point.y][point.x]
	(*matrix)[point.y][point.x] = (*matrix)[pointF3.y][pointF3.x]
	(*matrix)[pointF3.y][pointF3.x] = (*matrix)[pointF2.y][pointF2.x]
	(*matrix)[pointF2.y][pointF2.x] = (*matrix)[pointF1.y][pointF1.x]
	(*matrix)[pointF1.y][pointF1.x] = tmp
	fmt.Println(point, pointF1, pointF2, pointF3)
}

func rotate(matrix [][]int) {
	for i := range len(matrix) / 2 {
		for j := range (len(matrix) + 1) / 2 {
			rotateSingle(pos{j, i}, &matrix)
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	rotate(matrix)
	ans := matrix

	fmt.Println("\noutput:", Serialize(ans))
}
