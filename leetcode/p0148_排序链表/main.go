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

// 迭代，空间复杂度为常数
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	listLen := 0
	for node := head; node != nil; node = node.Next {
		listLen++
	}

	fakeHead := &ListNode{Next: head}
	for subLen := 1; subLen < listLen; subLen <<= 1 {
		prev, cur := fakeHead, fakeHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLen && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLen && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return fakeHead.Next
}
