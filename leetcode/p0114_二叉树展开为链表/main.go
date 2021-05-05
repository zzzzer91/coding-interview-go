// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 利用前序遍历
func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	var res []*TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			res = append(res, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}

// 递归
func flatten2(root *TreeNode) {
	var pre *TreeNode
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Right)
		f(root.Left)
		root.Right = pre
		root.Left = nil
		pre = root
	}
	f(root)
}
