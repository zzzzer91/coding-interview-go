// https://leetcode-cn.com/problems/sort-list/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归，但空间复杂度不满足常数级别
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	midNode := findMidNode(head)
	tmp := midNode.Next
	midNode.Next = nil
	return merge(sortList(head), sortList(tmp))
}

func findMidNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slowP, fastP := head, head.Next
	for fastP != nil && fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}
	return slowP
}

func merge(p, q *ListNode) *ListNode {
	fakeHead := ListNode{}
	cur := &fakeHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}
	if p != nil {
		cur.Next = p
	} else if q != nil {
		cur.Next = q
	}
	return fakeHead.Next
}
