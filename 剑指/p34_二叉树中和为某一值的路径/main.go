// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var comb []int
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root.Left == nil && root.Right == nil {
			if target == root.Val {
				temp := make([]int, len(comb), len(comb)+1)
				copy(temp, comb)
				res = append(res, append(temp, root.Val))
			}
			return
		}
		comb = append(comb, root.Val)
		if root.Left != nil {
			dfs(root.Left, target-root.Val)
		}
		if root.Right != nil {
			dfs(root.Right, target-root.Val)
		}
		comb = comb[:len(comb)-1]
	}
	dfs(root, sum)
	return res
}
