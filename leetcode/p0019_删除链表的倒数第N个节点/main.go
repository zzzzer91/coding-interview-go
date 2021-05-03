// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := ListNode{Next: head}
	slowP, fastP := &fakeHead, &fakeHead
	for i := 0; i < n; i++ {
		fastP = fastP.Next
	}
	for fastP.Next != nil {
		slowP, fastP = slowP.Next, fastP.Next
	}
	slowP.Next = slowP.Next.Next
	return fakeHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	slowP, fastP := head, head
	for i := 0; i < n; i++ {
		fastP = fastP.Next
	}
	if fastP == nil { // 只有1个节点的情况
		return head.Next
	}
	for fastP.Next != nil {
		slowP, fastP = slowP.Next, fastP.Next
	}
	slowP.Next = slowP.Next.Next
	return head
}
