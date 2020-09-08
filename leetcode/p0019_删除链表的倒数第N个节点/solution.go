package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := ListNode{Next: head}
	slowP := &fakeHead
	fastP := &fakeHead
	for i := 0; i < n; i++ {
		fastP = fastP.Next
		if fastP == nil {
			return nil
		}
	}
	for fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next
	}
	slowP.Next = slowP.Next.Next
	return fakeHead.Next
}
