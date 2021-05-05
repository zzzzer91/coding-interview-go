// https://leetcode-cn.com/problems/maximal-square/

package main

func main() {

}

func maximalSquare(matrix [][]byte) int {
	maxEdgeLen := 0
	m, n := len(matrix), len(matrix[0])
	// dp[i][j]表示以matrix[i - 1][j - 1]为右下角所能构成的最大正方形边长，
	// dp数组最左边一列和最上边一行都是0，可以去除一些条件判断。
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				// 仔细思考为什么这样是成立的，
				// 关键在取最小值
				// 只有三者都为 1（或者其他更大的边长）的时候，
				// 再加上 1才能组成一个更大的正方形，
				// 否则正方形边长只能还是 1
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
				maxEdgeLen = max(maxEdgeLen, dp[i][j])
			}
		}
	}
	return maxEdgeLen * maxEdgeLen // 不能直接返回 dp[m][n] * dp[m][n]，因为用了 min 而不是 max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
