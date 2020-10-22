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

func f(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && f(node1.Left, node2.Right) && f(node1.Right, node2.Left)
}
