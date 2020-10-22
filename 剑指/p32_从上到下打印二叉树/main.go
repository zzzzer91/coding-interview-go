// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		root = q.Remove(q.Front()).(*TreeNode)
		res = append(res, root.Val)
		if root.Left != nil {
			q.PushBack(root.Left)
		}
		if root.Right != nil {
			q.PushBack(root.Right)
		}
	}
	return res
}
