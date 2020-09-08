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
		fakeHead.Next, head.Next = head, fakeHead.Next
		head = temp
	}
	return fakeHead.Next
}
