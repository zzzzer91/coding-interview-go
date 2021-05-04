// https://leetcode-cn.com/problems/unique-paths/

package main

func main() {

}

// 还可以进一步把二维数组优化成一维数组
// m 是列数，n 是行数
func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		// 初始化初始值
		if i == 0 {
			for j := 0; j < m; j++ {
				dp[0][j] = 1
			}
		} else {
			dp[i][0] = 1
		}
	}
	// 状态转移
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[n-1][m-1]
}

func uniquePaths2(m int, n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			// dp[j] 实际上储存了上一行该列的值，所以二维数组可以省略成一维数组
			// 滚动数组
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
