// https://leetcode-cn.com/problems/reverse-linked-list/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	fakeHead := ListNode{}
	for head != nil {
		fakeHead.Next, head.Next, head = head, fakeHead.Next, head.Next
	}
	return fakeHead.Next
}
