// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(f(root.Left), 0)
		rightGain := max(f(root.Right), 0)
		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		res = max(res, root.Val+leftGain+rightGain)
		// 返回节点的最大贡献值
		// 两条路径只能选一条
		return root.Val + max(leftGain, rightGain)
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
