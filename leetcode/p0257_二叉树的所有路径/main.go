// https://leetcode-cn.com/problems/binary-tree-paths/

package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var comb []string
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			comb = append(comb, strconv.Itoa(root.Val))
			res = append(res, strings.Join(comb, "->"))
			comb = comb[:len(comb)-1]
			return
		}
		comb = append(comb, strconv.Itoa(root.Val))
		dfs(root.Left)
		dfs(root.Right)
		comb = comb[:len(comb)-1]
	}
	dfs(root)
	return res
}

func binaryTreePaths2(root *TreeNode) []string {
	var res []string
	var dfs func(root *TreeNode, s string)
	dfs = func(root *TreeNode, s string) {
		if root == nil {
			return
		}
		s += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			res = append(res, s)
		} else {
			s += "->"
			dfs(root.Left, s)
			dfs(root.Right, s)
		}
	}
	dfs(root, "")
	return res
}
