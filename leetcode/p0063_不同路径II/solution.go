package p0063_不同路径II

// 感觉还不够好
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rowNum := len(obstacleGrid)
	colNum := len(obstacleGrid[0])
	// 多分配一行和一列，下面状态转移时，可以免去一些额外判断
	dp := make([][]int, rowNum+1)
	for i := 0; i <= rowNum; i++ {
		dp[i] = make([]int, colNum+1)
	}
	// 注意这里 dp[0][1]=1 或 dp[1][0]=1 都行
	dp[0][1] = 1
	for i := 1; i <= rowNum; i++ {
		for j := 1; j <= colNum; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[rowNum][colNum]
}

// 还是可以优化成一维数组
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	colNum := len(obstacleGrid[0])
	dp := make([]int, colNum)
	dp[0] = 1
	for _, row := range obstacleGrid {
		for j, a := range row {
			if a == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[colNum-1]
}
