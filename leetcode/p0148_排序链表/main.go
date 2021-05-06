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

// 迭代，归并排序，空间复杂度为常数
func sortList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	listLen := 0
	for p := head; p != nil; p = p.Next {
		listLen++
	}
	fakeHead := ListNode{Next: head}
	for size := 1; size < listLen; size <<= 1 {
		cur := fakeHead.Next // 注意这里不能用 head，head 排完序可能就不再头部了
		tail := &fakeHead    // 用于将分割排序后的链表重新连接起来
		for cur != nil {
			left := cur                    // left(cur)->2->1->3->4->5...
			right := cut(left, size)       // left(cur)->2 right->1->3->4->5...
			cur = cut(right, size)         // left->2 right->1 cur->3->4->5->...
			tail.Next = merge(left, right) // left(tail.Next)->1->2 cur->3->4->5->...
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return fakeHead.Next
}

// 返回 head 跳过 size-1 个节点后的节点，并在节点前切断
func cut(head *ListNode, size int) *ListNode {
	p := head
	for p != nil && size > 1 {
		p = p.Next
		size--
	}
	if p == nil {
		return nil
	}
	next := p.Next
	p.Next = nil
	return next
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
