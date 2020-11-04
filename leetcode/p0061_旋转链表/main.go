// https://leetcode-cn.com/problems/rotate-list/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	n := 1
	p := head
	for ; p.Next != nil; p = p.Next {
		n++
	}
	p.Next = head
	n = n - k%n
	p = head
	for ; n > 1; n-- {
		p = p.Next
	}
	newHead := p.Next
	p.Next = nil
	return newHead
}
