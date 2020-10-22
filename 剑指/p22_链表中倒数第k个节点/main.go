// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slowP, fastP := head, head
	for i := 0; i < k; i++ {
		fastP = fastP.Next
	}
	for fastP != nil {
		slowP, fastP = slowP.Next, fastP.Next
	}
	return slowP
}
