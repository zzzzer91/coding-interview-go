// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/

package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	q := list.New()
	q.PushBack(root)
	flag := true
	for q.Len() > 0 {
		var row []int
		for i := q.Len(); i > 0; i-- {
			var node *TreeNode
			if flag {
				node = q.Remove(q.Front()).(*TreeNode)
				if node.Left != nil {
					q.PushBack(node.Left)
				}
				if node.Right != nil {
					q.PushBack(node.Right)
				}
			} else {
				node = q.Remove(q.Back()).(*TreeNode)
				if node.Right != nil {
					q.PushFront(node.Right)
				}
				if node.Left != nil {
					q.PushFront(node.Left)
				}
			}
			row = append(row, node.Val)
		}
		flag = !flag
		res = append(res, row)
	}
	return res
}
