// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 二叉搜索树的中序遍历可以得到一个有序升序序列
// 如果，先搜索二叉树右子树，这样得到的结果是降序的
func kthLargest(root *TreeNode, k int) int {
	res := 0
	count := 0
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = root.Val
			count++
			if count == k {
				break
			}
			root = root.Left
		}
	}
	return res
}
