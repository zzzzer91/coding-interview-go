package p0445_两数相加II

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	fakeHead := ListNode{}
	val := 0
	for len(stack1) != 0 || len(stack2) != 0 || val != 0 {
		if len(stack1) != 0 {
			val += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			val += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		// 插到首节点
		fakeHead.Next = &ListNode{Val: val % 10, Next: fakeHead.Next}
		val /= 10
	}
	return fakeHead.Next
}
