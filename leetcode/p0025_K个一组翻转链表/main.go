// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println(reverseKGroup(head, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	fakeHead := ListNode{}
	slowP := &fakeHead
	fastP := head // 探测剩余元素是否满足数量要求，不满足不需要反转
	count := 0
	var saveNode *ListNode
	for head != nil {
		if count%k == 0 {
			saveNode = head // 记录头节点，该节点翻转后会变成最后一个节点
			for i := 1; i < k; i++ {
				fastP = fastP.Next
				if fastP == nil {
					slowP.Next = head // 剩余元素不满足数量要求，不需要反转
					return fakeHead.Next
				}
			}
			fastP = fastP.Next
		}
		slowP.Next, head, head.Next = head, head.Next, slowP.Next
		count++
		if count%k == 0 {
			slowP = saveNode
		}
	}
	return fakeHead.Next
}

// 更好
func reverseKGroup2(head *ListNode, k int) *ListNode {
	fakeHead := ListNode{}
	fakeHead.Next = head
	pre := &fakeHead
	fastP := pre
	for fastP != nil {
		// 找到反转的范围
		for i := 0; i < k && fastP != nil; i++ {
			fastP = fastP.Next
		}
		if fastP == nil {
			break
		}

		next := fastP.Next
		fastP.Next = nil
		// 链表反转,范围[cur, fastP]
		cur := pre.Next
		pre.Next = reverseLink(cur)
		// 前变后
		pre = cur
		cur.Next = next
		fastP = pre // 回退一个节点
	}
	return fakeHead.Next
}

func reverseLink(head *ListNode) *ListNode {
	fakeHead := ListNode{}
	p := &fakeHead
	for p != nil {
		p.Next, head, head.Next = head, head.Next, p.Next
	}
	return fakeHead.Next
}
