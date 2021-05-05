// https://leetcode-cn.com/problems/palindrome-linked-list/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 时 O(n)，空 O(1)，思想：
// 使用快慢指针（快走两步，慢走一步），当快指针走到底后，慢指针正好到一半，
// 然后逆置后一半链表，然后比较。如果不能修改原链表，那么结束后再逆置回去。
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 当链表长度为奇数，slowP 指向最中间那个节点
	slowP, fastP := head, head
	for fastP.Next != nil && fastP.Next.Next != nil {
		slowP, fastP = slowP.Next, fastP.Next.Next
	}
	p1, p2 := head, reverse(slowP.Next)
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1, p2 = p1.Next, p2.Next
	}
	reverse(slowP.Next)
	return true
}

func reverse(head *ListNode) *ListNode {
	fakeHead := ListNode{}
	for head != nil {
		fakeHead.Next, head, head.Next = head, head.Next, fakeHead.Next
	}
	return fakeHead.Next
}
