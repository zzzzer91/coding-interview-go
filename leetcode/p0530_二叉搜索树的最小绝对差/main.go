// https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	res := int(1e9)
	var stack []*TreeNode
	pre := -1
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != -1 && root.Val-pre < res {
				res = root.Val - pre
			}
			pre = root.Val
			root = root.Right
		}
	}
	return res
}
