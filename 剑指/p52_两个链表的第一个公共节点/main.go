// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	listALen, listBLen := 0, 0
	pA, pB := headA, headB
	for pA != nil {
		listALen++
		pA = pA.Next
	}
	for pB != nil {
		listBLen++
		pB = pB.Next
	}
	pA, pB = headA, headB
	if listALen > listBLen {
		n := listALen - listBLen
		for n > 0 {
			pA = pA.Next
			n--
		}
	} else {
		n := listBLen - listALen
		for n > 0 {
			pB = pB.Next
			n--
		}
	}
	for pA != nil {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}
	return nil
}

// 非常优美的解法
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil { // p1没有了接headB
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil { // p2没有了接headA
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1 // 即使有不相交的特殊情况，此时 p1 和 p2 都指向 nil
}
