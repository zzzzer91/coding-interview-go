package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNode(v int) *TreeNode {
	return &TreeNode{Val: v}
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}

// 后序遍历相比较另外两种遍历需要一个额外变量记录回溯节点
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var prev *TreeNode // 记录从哪一个节点回溯过来的
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			if root.Right == nil || root.Right == prev {
				stack = stack[:len(stack)-1]
				res = append(res, root.Val)
				prev = root
				root = nil
			} else {
				root = root.Right
			}
		}
	}
	return res
}

func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	queue := list.New()
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
		ret = append(ret, level)
	}
	return ret
}

func main() {
	/*
	 * 创建一棵树
	 *
	 *      1
	 *     / \
	 *    2   5
	 *     \
	 *      3
	 *     /
	 *    4
	 */
	root := CreateNode(1)
	root.Left = CreateNode(2)
	root.Left.Right = CreateNode(3)
	root.Left.Right.Left = CreateNode(4)
	root.Left.Right.Right = CreateNode(5)
	root.Right = CreateNode(6)
	root.Right.Left = CreateNode(7)

	// 先序遍历
	fmt.Println(preorderTraversal(root))
	// 中序遍历
	fmt.Println(inorderTraversal(root))
	// 后序遍历
	fmt.Println(postorderTraversal(root))
	// 层序遍历
	fmt.Println(levelOrder(root))
}
