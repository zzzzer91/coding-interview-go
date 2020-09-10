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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
