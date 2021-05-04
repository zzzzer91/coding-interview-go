// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var level []int // 每个循环 level 都会被置 nil
		for i := queue.Len(); i > 0; i-- {
			node := queue.Remove(queue.Front()).(*TreeNode)
			level = append(level, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
