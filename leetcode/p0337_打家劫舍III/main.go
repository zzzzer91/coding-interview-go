// https://leetcode-cn.com/problems/house-robber-iii/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

// res[0] 代表选了当前节点的最大值，res[1] 代表没选当前节点的最大值
func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
