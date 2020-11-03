// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			count++
			if count == k {
				return root.Val
			}
			root = root.Right
		}
	}
	return -1
}
