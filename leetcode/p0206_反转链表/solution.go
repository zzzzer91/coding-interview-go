package p0206_反转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	fakeHead := ListNode{}
	for head != nil {
		temp := head.Next
		fakeHead.Next, head.Next = head, fakeHead.Next
		head = temp
	}
	return fakeHead.Next
}
