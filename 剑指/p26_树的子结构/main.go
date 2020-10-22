// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return f(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func f(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return A.Val == B.Val && f(A.Left, B.Left) && f(A.Right, B.Right)
}
