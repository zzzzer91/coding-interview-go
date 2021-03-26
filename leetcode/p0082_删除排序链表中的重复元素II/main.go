// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	fakeHead := ListNode{Next: head}
	pre := &fakeHead
	for pre.Next != nil {
		var p2 *ListNode
		for p2 = pre.Next.Next; p2 != nil && p2.Val == pre.Next.Val; p2 = p2.Next {
		}
		if p2 != pre.Next.Next {
			pre.Next = p2
		} else {
			pre = pre.Next
		}
	}
	return fakeHead.Next
}
