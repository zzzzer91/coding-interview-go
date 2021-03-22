// https://leetcode-cn.com/problems/set-matrix-zeroes/

package main

func main() {

}

// 常规思路
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rowLen, colLen := len(matrix), len(matrix[0])
	var points [][2]int
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if matrix[i][j] == 0 {
				points = append(points, [2]int{i, j})
			}
		}
	}
	xSet, ySet := make(map[int]bool), make(map[int]bool)
	for _, p := range points {
		if !xSet[p[0]] {
			for i := 0; i < colLen; i++ {
				matrix[p[0]][i] = 0
			}
			xSet[p[0]] = true
		}
		if !ySet[p[1]] {
			for i := 0; i < rowLen; i++ {
				matrix[i][p[1]] = 0
			}
			ySet[p[1]] = true
		}
	}
}

// 只用常量空间
func setZeroes2(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rowLen, colLen := len(matrix), len(matrix[0])

	// 判断第一行和第一列是否有0出现
	rowFlag, colFlag := false, false
	for i := 0; i < rowLen; i++ {
		if matrix[i][0] == 0 {
			rowFlag = true
		}
	}
	for i := 0; i < colLen; i++ {
		if matrix[0][i] == 0 {
			colFlag = true
		}
	}

	// 从 matrix[1][1]开始迭代，如果出现零，则将该位置对应的第一行和第一列置为0
	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 判断第一行和第一列是否需要全部置0
	if rowFlag {
		for i := 0; i < rowLen; i++ {
			matrix[i][0] = 0
		}
	}
	if colFlag {
		for i := 0; i < colLen; i++ {
			matrix[0][i] = 0
		}
	}
}
