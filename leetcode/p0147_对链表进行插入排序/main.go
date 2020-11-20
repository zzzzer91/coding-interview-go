// https://leetcode-cn.com/problems/insertion-sort-list/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val { // 确保 lastSorted 及之前都有序，然后就可以找位置插入
			lastSorted = lastSorted.Next
		} else { // 从有序链表中找第一个比当前大的节点，然后在其之前插入
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
