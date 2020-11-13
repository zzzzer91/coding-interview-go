package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

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

// 非常优美的解法
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/solution/tu-jie-xiang-jiao-lian-biao-by-user7208t/
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	// p1 和 p2 走了同样长的路程，所以一定同时到达终点
	for p1 != p2 {
		if p1 == nil { // p1 没有了接 headB
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil { // p2 没有了接 headA
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1 // 即使有不相交的特殊情况，此时 p1 和 p2 都指向 nil
}
