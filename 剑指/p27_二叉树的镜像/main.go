// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 注意：
	// root.Left = mirrorTree(root.Right)
	// root.Right = mirrorTree(root.Left)
	// 这样是不对的，Left 的值被改变了，传入 Right 就错了
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}
