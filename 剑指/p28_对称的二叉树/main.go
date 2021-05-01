// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return f(root.Left, root.Right)
}

func f(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && f(l.Left, r.Right) && f(l.Right, r.Left)
}
