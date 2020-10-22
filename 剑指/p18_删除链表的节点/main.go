// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	fakeHead := ListNode{Next: head}
	pre := &fakeHead
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
			break
		}
		pre = pre.Next
	}
	return fakeHead.Next
}

// 递归
func deleteNode2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = deleteNode(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
