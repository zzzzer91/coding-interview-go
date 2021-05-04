// https://leetcode-cn.com/problems/symmetric-tree/
// 判断二叉树是否对称
/*
    1
   / \
  2   2
 / \ / \
3  4 4  3
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return f(root.Left, root.Right)
}

func f(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && f(p.Left, q.Right) && f(p.Right, q.Left)
}
