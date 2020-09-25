// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 {
		return node
	}
	rootIndex := 0
	for node.Val != inorder[rootIndex] {
		rootIndex++
	}
	node.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	node.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	return node
}
