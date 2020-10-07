// https://leetcode-cn.com/problems/validate-binary-search-tree/

package main

import "fmt"

func main() {
	root := CreateNode(5)
	root.Left = CreateNode(1)
	root.Right = CreateNode(4)
	root.Right.Left = CreateNode(3)
	root.Right.Right = CreateNode(6)
	fmt.Println(isValidBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNode(v int) *TreeNode {
	return &TreeNode{Val: v}
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && root.Val <= pre.Val {
				return false
			}
			pre = root
			root = root.Right
		}
	}
	return true
}
