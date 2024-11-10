// Created by OrkWard at 2024/11/09 11:44
// leetgo: dev
// https://leetcode.com/problems/rotate-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func listLen(head *ListNode) (length int) {
	p := head
	for p != nil {
		p = p.Next
		length += 1
	}
	return
}

func listLast(head *ListNode) (last *ListNode) {
	last = head
	for last.Next != nil {
		last = last.Next
	}
	return
}

func rotateRight(head *ListNode, k int) (ans *ListNode) {
	if (head == nil) {
		return nil
	}

	print("Length:", listLen(head))

	// new head index in old list
	length := listLen(head)
	startIdx := (length - k % length) % length

	if (startIdx == 0) {
		return head
	}

	// new last index in old list
	lastIdx := startIdx - 1
	pNewLast := head
	for i := 1; i <= lastIdx; i += 1 {
		pNewLast = pNewLast.Next
	}

	ans = pNewLast.Next
	pLast := listLast(ans)
	pLast.Next = head
	pNewLast.Next = nil
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := rotateRight(head, k)

	fmt.Println("\noutput:", Serialize(ans))
}
