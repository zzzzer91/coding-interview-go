// https://leetcode-cn.com/problems/reverse-linked-list-ii/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	fakeHead := ListNode{Next: head}
	pre := &fakeHead
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	head = pre.Next
	for i := m; i < n; i++ {
		tmp := head.Next
		head.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}
	return fakeHead.Next
}
