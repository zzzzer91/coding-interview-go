// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(r *TreeNode, sum int)
	dfs = func(r *TreeNode, sum int) {
		if r == nil {
			return
		}
		sum = sum*10 + r.Val
		if r.Left == nil && r.Right == nil {
			res += sum
		} else {
			dfs(r.Left, sum)
			dfs(r.Right, sum)
		}
	}
	dfs(root, 0)
	return res
}
