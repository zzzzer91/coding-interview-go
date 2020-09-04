package p0543_二叉树的直径

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 效率不高，有重复计算
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 求三者中最大
	return max(max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)), maxDepth(root.Left)+maxDepth(root.Right))
}

func diameterOfBinaryTree2(root *TreeNode) int {
	ans := 0
	getDeepestDepth(root, &ans)
	return ans
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func getDeepestDepth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	l := getDeepestDepth(root.Left, ans)
	r := getDeepestDepth(root.Right, ans)
	*ans = max(*ans, l+r)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
