// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 有两种做法：1、堆（优先级队列）2、归并

package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func main() {
}

// 堆
func mergeKLists(lists []*ListNode) *ListNode {
	var h MinHeap
	for _, l := range lists {
		if l != nil {
			heap.Push(&h, l)
		}
	}
	fakeHead := ListNode{}
	p := &fakeHead
	for h.Len() > 0 {
		node := heap.Pop(&h).(*ListNode)
		p.Next = node
		p = p.Next
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
	}
	return fakeHead.Next
}

// 归并，牛逼
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for gap := 1; gap < len(lists); gap *= 2 {
		for i := 0; i+gap < len(lists); i += 2 * gap {
			lists[i] = merge(lists[i], lists[i+gap])
		}
	}
	return lists[0]
}

func merge(l1, l2 *ListNode) *ListNode {
	list := &ListNode{}
	visitNode := list
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			visitNode.Next = l1
			visitNode = visitNode.Next
			l1 = l1.Next
		} else {
			visitNode.Next = l2
			visitNode = visitNode.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		visitNode.Next = l1
	} else {
		visitNode.Next = l2
	}
	return list.Next
}
