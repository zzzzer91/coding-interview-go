package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isBalanced(root *TreeNode) bool {
	return depth(root) != -1
}

// 在二叉树平衡的情况下返回树的高度，否则返回 -1
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := depth(root.Left)
	if lDepth < 0 {
		return -1
	}
	rDepth := depth(root.Right)
	if rDepth < 0 {
		return -1
	}
	if abs(lDepth-rDepth) > 1 {
		return -1
	}
	return max(lDepth, rDepth) + 1
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
