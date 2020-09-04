package p0021_合并两个有序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	curNode := &head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curNode.Next = l1
			l1 = l1.Next
		} else {
			curNode.Next = l2
			l2 = l2.Next
		}
		curNode = curNode.Next
	}
	if l1 != nil {
		curNode.Next = l1
	} else {
		curNode.Next = l2
	}

	return head.Next
}

// 递归
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2
}
