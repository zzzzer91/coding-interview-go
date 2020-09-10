package p0107_二叉树的层次遍历II

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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

	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}
