// https://leetcode-cn.com/problems/rotate-image/

package main

func main() {

}

func rotate(matrix [][]int) {
	n := len(matrix)
	// 先进行对角（左上到右下这条对角线）翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 再进行沿垂直中线左右翻转
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
