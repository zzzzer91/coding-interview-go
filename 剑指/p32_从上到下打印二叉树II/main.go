// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/

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
	for q.Len() > 0 {
		var temp []int
		for i := q.Len(); i > 0; i-- {
			root = q.Remove(q.Front()).(*TreeNode)
			temp = append(temp, root.Val)
			if root.Left != nil {
				q.PushBack(root.Left)
			}
			if root.Right != nil {
				q.PushBack(root.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
