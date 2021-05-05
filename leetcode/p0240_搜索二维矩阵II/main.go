// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/

package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1 // 从右上角看可以发现是一棵二叉搜索树
	for x < m && y >= 0 {
		if matrix[x][y] > target { // 搜二叉树左节点
			y--
		} else if matrix[x][y] < target { // 搜二叉树右节点
			x++
		} else {
			return true
		}
	}
	return false
}
