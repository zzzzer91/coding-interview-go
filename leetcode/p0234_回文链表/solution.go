package p0234_回文链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	fakeHead := ListNode{}
	for head != nil {
		temp := head.Next
		head.Next = fakeHead.Next
		fakeHead.Next = head
		head = temp
	}
	return fakeHead.Next
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
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}

	p := reverse(slowP.Next)
	for p != nil {
		if head.Val != p.Val {
			return false
		}
		head = head.Next
		p = p.Next
	}
	return true
}
