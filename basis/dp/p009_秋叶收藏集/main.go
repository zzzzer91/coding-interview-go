// https://leetcode-cn.com/problems/UlBDOe/

package main

func main() {

}

func minimumOperations(leaves string) int {
	// dp[0] 代表前面全是红叶的情况
	// dp[1] 代表前面是红叶、黄叶的情况
	// dp[2] 代表前面是红叶、黄叶、红叶的情况
	var dp [3][]int
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, len(leaves))
	}
	for i := 0; i < len(leaves); i++ {
		if i == 0 {
			dp[0][i] = boolToInt(leaves[i] != 'r')
		} else {
			dp[0][i] = dp[0][i-1] + boolToInt(leaves[i] != 'r')
		}
		if i == 0 {
			dp[1][i] = dp[0][i]
		} else {
			dp[1][i] = min(dp[0][i-1], dp[1][i-1]) + boolToInt(leaves[i] != 'y')
		}
		if i < 2 {
			dp[2][i] = dp[1][i]
		} else {
			dp[2][i] = min(dp[1][i-1], dp[2][i-1]) + boolToInt(leaves[i] != 'r')
		}
	}
	return dp[2][len(leaves)-1]
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
