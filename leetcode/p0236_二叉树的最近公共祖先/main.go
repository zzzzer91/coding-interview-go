// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

/**
  注意p,q必然存在树内, 且所有节点的值唯一!!!
  递归思想, 对以root为根的(子)树进行查找p和q, 如果root == null || p || q 直接返回root
  表示对于当前树的查找已经完毕, 否则对左右子树进行查找, 根据左右子树的返回值判断:
  1. 左右子树的返回值都不为null, 由于值唯一左右子树的返回值就是p和q, 此时root为LCA
  2. 如果左右子树返回值只有一个不为null, 说明只有p和q存在与左或右子树中, 最先找到的那个节点为LCA
  3. 左右子树返回值均为null, p和q均不在树中, 返回null
  **/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}
