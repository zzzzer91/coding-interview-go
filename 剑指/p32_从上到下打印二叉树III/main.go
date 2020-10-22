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
	flag := false // flag 为 true，则反向打印
	for q.Len() > 0 {
		var temp []int
		if flag {
			for i := q.Len(); i > 0; i-- {
				root = q.Remove(q.Back()).(*TreeNode)
				temp = append(temp, root.Val)
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
				temp = append(temp, root.Val)
				if root.Left != nil {
					q.PushBack(root.Left)
				}
				if root.Right != nil {
					q.PushBack(root.Right)
				}
			}
		}
		flag = !flag
		res = append(res, temp)
	}
	return res
}
