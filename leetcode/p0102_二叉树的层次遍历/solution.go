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
	queue := list.New()
	var result [][]int
	queue.PushBack(root)
	for queue.Len() > 0 {
		// 每个循环 level 都会被置 nil
		var level []int
		for i := queue.Len(); i > 0; i-- {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			level = append(level, node.Val)
		}
		result = append(result, level)
	}
	return result
}
