// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val == root.Val {
		return p
	}
	if q.Val == root.Val {
		return q
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, q, p)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, q, p)
	} else {
		return root
	}
}

// 优化，使用迭代
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val { // 保证 p.val < q.val
		p, q = q, p
	}
	for root != nil {
		if root.Val < p.Val {
			root = root.Right
		} else if root.Val > q.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}
