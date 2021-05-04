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

func rotate2(matrix [][]int) {
	n := len(matrix)
	// 先沿水平中线上下翻转
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	// 再沿左上到右下的斜线翻转
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
