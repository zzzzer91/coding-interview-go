package p0203_移除链表元素

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	fakeHead := ListNode{}
	for head != nil {
		temp := head.Next
		head.Next = fakeHead.Next
		fakeHead.Next = head
		head = temp
	}

	return fakeHead.Next
}
