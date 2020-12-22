// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	q := list.New()
	q.PushBack(root)
	flag := false
	for q.Len() > 0 {
		row := make([]int, 0, q.Len())
		if flag {
			for i := q.Len(); i > 0; i-- {
				root = q.Remove(q.Back()).(*TreeNode)
				row = append(row, root.Val)
				if root.Right != nil {
					q.PushFront(root.Right)
				}
				if root.Left != nil {
					q.PushFront(root.Left)
				}
			}
		} else {
			for i := q.Len(); i > 0; i-- {
				root = q.Remove(q.Front()).(*TreeNode)
				row = append(row, root.Val)
				if root.Left != nil {
					q.PushBack(root.Left)
				}
				if root.Right != nil {
					q.PushBack(root.Right)
				}
			}
		}
		flag = !flag
		res = append(res, row)
	}
	return res
}
