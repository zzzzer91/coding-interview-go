// https://leetcode-cn.com/problems/sum-of-left-leaves/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
