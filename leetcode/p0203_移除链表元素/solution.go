package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	fakeHead := ListNode{}
	for head != nil {
		temp := head.Next
		head.Next = fakeHead.Next
		fakeHead.Next = head
		head = temp
	}

	return fakeHead.Next
}
