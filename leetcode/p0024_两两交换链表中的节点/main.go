// https://leetcode-cn.com/problems/swap-nodes-in-pairs/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	fakeHead := ListNode{}
	p := &fakeHead
	count := 0
	for head != nil {
		p.Next, head, head.Next = head, head.Next, p.Next
		count++
		if count&1 == 0 {
			p = p.Next.Next
		}
	}
	return fakeHead.Next
}
