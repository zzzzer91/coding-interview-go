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
	// 出现空树，匹配失败
	if A == nil || B == nil {
		return false
	}
	return f(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func f(A *TreeNode, B *TreeNode) bool {
	// 当节点 B 为空：说明树 B 已匹配完成
	if B == nil {
		return true
	}
	// 当节点 A 为空：说明已经越过树 A 叶子节点，即匹配失败
	if A == nil {
		return false
	}
	return A.Val == B.Val && f(A.Left, B.Left) && f(A.Right, B.Right)
}
