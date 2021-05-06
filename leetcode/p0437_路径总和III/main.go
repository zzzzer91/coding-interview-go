// https://leetcode-cn.com/problems/path-sum-iii/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 双重递归
func pathSum(root *TreeNode, sum int) int {
	res := 0
	var count func(r *TreeNode, target int)
	var dfs func(r *TreeNode)
	count = func(r *TreeNode, target int) {
		if r == nil {
			return
		}
		target -= r.Val
		if target == 0 {
			res++
		}
		count(r.Left, target)
		count(r.Right, target)
	}
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		count(r, sum)
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return res
}
