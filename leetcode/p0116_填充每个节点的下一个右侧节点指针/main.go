// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil { // 利用 next 指针
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

// 更易懂
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)
	connectTwoNode(node1.Right, node2.Left)
}
