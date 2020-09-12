// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/submissions/

package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		curSums := 0
		curCount := 0
		for i := q.Len(); i > 0; i-- {
			node := q.Remove(q.Front()).(*TreeNode)
			curSums += node.Val
			curCount++
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		res = append(res, float64(curSums)/float64(curCount))
	}
	return res
}
