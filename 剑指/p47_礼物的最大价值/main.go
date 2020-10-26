// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/

package main

func main() {

}

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rowLen, colLen := len(grid), len(grid[0])
	dp := make([]int, colLen+1)
	for i := 0; i < rowLen; i++ {
		for j := 1; j <= colLen; j++ {
			dp[j] = max(dp[j], dp[j-1]) + grid[i][j-1]
		}
	}
	return dp[colLen]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 节约空间
func maxValue2(grid [][]int) int {
	rowLen := len(grid)
	colLen := len(grid[0])
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[0][j] += grid[0][j-1]
			} else if j == 0 {
				grid[i][0] += grid[i-1][0]
			} else {
				grid[i][j] = max(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	return grid[rowLen-1][colLen-1]
}
