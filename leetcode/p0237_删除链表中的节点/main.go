// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
