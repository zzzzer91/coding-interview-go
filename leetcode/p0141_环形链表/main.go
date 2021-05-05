// https://leetcode-cn.com/problems/linked-list-cycle/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slowP, fastP := head, head
	for fastP.Next != nil && fastP.Next.Next != nil {
		slowP, fastP = slowP.Next, fastP.Next.Next
		if slowP == fastP {
			return true
		}
	}
	return false
}
