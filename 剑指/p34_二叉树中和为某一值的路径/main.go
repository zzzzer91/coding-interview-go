// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var comb []int
	var f func(root *TreeNode, target int)
	f = func(root *TreeNode, target int) {
		comb = append(comb, root.Val)
		target -= root.Val
		if target == 0 && (root.Left == nil && root.Right == nil) {
			res = append(res, append(make([]int, 0, len(comb)), comb...))
		} else {
			if root.Left != nil {
				f(root.Left, target)
			}
			if root.Right != nil {
				f(root.Right, target)
			}
		}
		comb = comb[:len(comb)-1]
	}
	f(root, target)
	return res
}
