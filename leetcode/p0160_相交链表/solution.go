package p0160_相交链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenListA := 0
	for p := headA; p != nil; p = p.Next {
		lenListA++
	}
	lenListB := 0
	for p := headB; p != nil; p = p.Next {
		lenListB++
	}
	sub := lenListA - lenListB
	if sub < 0 {
		for sub < 0 {
			headB = headB.Next
			sub++
		}
	} else {
		for sub > 0 {
			headA = headA.Next
			sub--
		}
	}
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
