// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/

package main

func main() {

}

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = max(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1]
}

func maxValue2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n+1) // 多开辟一格，代码写起来简洁很多
	for i := 0; i < m; i++ {
		for j := 1; j <= n; j++ {
			dp[j] = max(dp[j], dp[j-1]) + grid[i][j-1]
		}
	}
	return dp[n]
}

// 节约空间
func maxValue3(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] = max(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
