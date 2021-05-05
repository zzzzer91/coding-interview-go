// https://leetcode-cn.com/problems/linked-list-cycle-ii/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slowP, fastP := head, head
	for fastP.Next != nil && fastP.Next.Next != nil {
		slowP, fastP = slowP.Next, fastP.Next.Next
		if slowP == fastP {
			slowP = head
			for slowP != fastP {
				slowP, fastP = slowP.Next, fastP.Next
			}
			return slowP
		}
	}
	return nil
}
