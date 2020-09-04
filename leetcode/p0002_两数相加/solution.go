package p0002_两数相加

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeHead := ListNode{}
	p := &fakeHead
	carry := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		if val >= 10 {
			val -= 10
			carry = 1
		} else {
			carry = 0
		}
		p.Next = &ListNode{Val: val}
		l1 = l1.Next
		l2 = l2.Next
		p = p.Next
	}
	// 处理剩余节点
	for l1 != nil {
		val := l1.Val + carry
		if val >= 10 {
			val -= 10
			carry = 1
		} else {
			carry = 0
		}
		p.Next = &ListNode{Val: val}
		l1 = l1.Next
		p = p.Next

	}
	for l2 != nil {
		val := l2.Val + carry
		if val >= 10 {
			val -= 10
			carry = 1
		} else {
			carry = 0
		}
		p.Next = &ListNode{Val: val}
		l2 = l2.Next
		p = p.Next
	}

	if carry == 1 {
		p.Next = &ListNode{Val: 1}
	}

	return fakeHead.Next
}

// 更简洁，但性能差点
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeHead := ListNode{}
	p := &fakeHead
	val := 0
	for l1 != nil || l2 != nil || val != 0 {
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{Val: val % 10}
		val /= 10
		p = p.Next
	}

	return fakeHead.Next
}
