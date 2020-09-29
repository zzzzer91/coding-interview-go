// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Right == nil || root.Right == prev {
				res = append(res, root.Val)
				prev = root
				root = nil
			} else {
				stack = append(stack, root)
				root = root.Right
			}
		}
	}
	return res
}
