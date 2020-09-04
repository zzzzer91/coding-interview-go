// 判断二叉树是否对称

package p0101_symmetric__tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
    1
   / \
  2   2
 / \ / \
3  4 4  3
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil || q == nil) || p.Val != q.Val {
		return false
	}
	return isMirror(p.Left, q.Right) && isMirror(q.Left, p.Right)
}
