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
		slowP = slowP.Next
		fastP = fastP.Next.Next
		if slowP == fastP {
			return true
		}
	}
	return false
}
