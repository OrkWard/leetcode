// Created by OrkWard at 2024/11/21 15:28
// leetgo: 1.4.10
// https://leetcode.com/problems/simplify-path/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type Path struct {
	name string
	next *Path
	prev *Path
}

func parsePath(path string) *Path {
	p := path
	head := &Path{"/", nil, nil}
	head.prev = head // root.prev == root
	tail := head
	for len(p) > 1 {
		nextSlash := strings.Index(p[1:], "/") + 1
		// println("next: ", nextSlash)
		var name string

		if nextSlash-1 == -1 {
			name = p[1:]
			p = ""
		} else {
			name = p[1:nextSlash]
			p = p[nextSlash:]
		}

		// println("name: ", name)

		if name == "." || name == "" {
			continue
		} else if name == ".." {
			if tail != tail.prev {
				tail = tail.prev
				tail.next = nil
			}
		} else {
			tail.next = &Path{name, nil, nil}
			tail.next.prev = tail
			tail = tail.next
		}

		// println("tail: ", tail.name, tail)
	}

	return head
}

func genPathString(path *Path) (ans string) {
	if path.next == nil {
		return "/"
	}

	p := path.next
	for p != nil {
		ans += ("/" + p.name)
		p = p.next
	}
	return
}

func simplifyPath(path string) string {
	p := parsePath(path)
	return genPathString(p)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	path := Deserialize[string](ReadLine(stdin))
	ans := simplifyPath(path)

	fmt.Println("\noutput:", Serialize(ans))
}
