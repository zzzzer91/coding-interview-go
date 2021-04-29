// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var res []int
	var f func(head *ListNode)
	f = func(head *ListNode) {
		if head == nil {
			return
		}
		f(head.Next)
		res = append(res, head.Val)
	}
	f(head)
	return res
}

// 简洁写法，很妙
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)
}
