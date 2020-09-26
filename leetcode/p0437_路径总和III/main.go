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
	var count func(r *TreeNode, tmp int)
	var dfs func(r *TreeNode)
	count = func(r *TreeNode, tmp int) {
		if r == nil {
			return
		}
		tmp -= r.Val
		if tmp == 0 {
			res++
		}
		count(r.Left, tmp)
		count(r.Right, tmp)
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
