// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return f(root) != -1
}

func f(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := f(root.Left)
	if l == -1 {
		return -1
	}
	r := f(root.Right)
	if r == -1 {
		return -1
	}
	if abs(l-r) > 1 {
		return -1
	}
	return max(l, r) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
