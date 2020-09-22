// https://leetcode-cn.com/problems/binary-tree-cameras/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 0}
	fmt.Println(minCameraCover(root))
}

const (
	hasCamera  int = 0
	hasCovered int = 1
	notCovered int = 2
)

func minCameraCover(root *TreeNode) int {
	var res int
	if minCameraCoverInternal(root, &res) == notCovered {
		res++
	}
	return res
}

func minCameraCoverInternal(root *TreeNode, res *int) int {
	if root == nil {
		return hasCovered
	}
	left := minCameraCoverInternal(root.Left, res)
	right := minCameraCoverInternal(root.Right, res)
	if left == notCovered || right == notCovered {
		*res = *res + 1
		return hasCamera
	}
	if left == hasCamera || right == hasCamera {
		return hasCovered
	}
	return notCovered
}
