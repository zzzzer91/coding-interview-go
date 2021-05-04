// https://leetcode-cn.com/problems/minimum-path-sum/

package main

func main() {

}

func minPathSum(grid [][]int) int {
	colNum := len(grid[0])
	dp := make([]int, colNum)
	// 首先初始化第一行
	dp[0] = grid[0][0]
	for i := 1; i < colNum; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for _, row := range grid[1:] {
		for i, num := range row {
			if i == 0 {
				dp[i] += num
			} else {
				dp[i] = min(dp[i-1], dp[i]) + num
			}
		}
	}
	return dp[colNum-1]
}

// 把初始化第一行放进循环里，会多些判断，但代码更简洁
func minPathSum2(grid [][]int) int {
	colNum := len(grid[0])
	dp := make([]int, colNum)
	for i, row := range grid {
		for j, num := range row {
			if j == 0 {
				dp[j] += num
			} else if i == 0 {
				dp[j] = dp[j-1] + num
			} else {
				dp[j] = min(dp[j-1], dp[j]) + num
			}
		}
	}
	return dp[colNum-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
