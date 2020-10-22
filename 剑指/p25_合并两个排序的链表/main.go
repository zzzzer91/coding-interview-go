// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

// 非递归
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeHead := ListNode{}
	p := &fakeHead
	for l1 != nil || l2 != nil {
		if l1 == nil || (l2 != nil && l1.Val > l2.Val) {
			p.Next = l2
			l2 = l2.Next
		} else if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}
	return fakeHead.Next
}
