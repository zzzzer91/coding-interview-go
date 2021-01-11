// https://leetcode-cn.com/problems/triangle/

package main

func main() {

}

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	copy(dp, triangle[len(triangle)-1])
	for i := len(triangle) - 2; i >= 0; i-- {
		row := triangle[i]
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + row[j]
		}
	}
	return dp[0]
}

// 简化，不需要额外空间
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	for i := n - 2; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
