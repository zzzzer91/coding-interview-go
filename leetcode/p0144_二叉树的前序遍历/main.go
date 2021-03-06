package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for len(stack) != 0 || root != nil {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}
