// https://leetcode-cn.com/problems/diameter-of-binary-tree/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lc := f(root.Left)
		rc := f(root.Right)
		res = max(res, lc+rc) // 注意这里不需要 +1，具体看图可知
		return max(lc, rc) + 1
	}
	f(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
