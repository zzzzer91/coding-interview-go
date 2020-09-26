// https://leetcode-cn.com/problems/path-sum-ii/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var comb []int
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		if (root.Left == nil && root.Right == nil) && sum == root.Val {
			temp := make([]int, len(comb))
			copy(temp, comb)
			temp = append(temp, root.Val)
			res = append(res, temp)
			return
		}
		comb = append(comb, root.Val)
		dfs(root.Left, sum-root.Val)
		dfs(root.Right, sum-root.Val)
		comb = comb[:len(comb)-1]
	}
	dfs(root, sum)
	return res
}
