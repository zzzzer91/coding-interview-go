// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	fakeHead := ListNode{}
	for head != nil {
		fakeHead.Next, head, head.Next = head, head.Next, fakeHead.Next
	}
	return fakeHead.Next
}
