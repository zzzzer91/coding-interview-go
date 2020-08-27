// https://www.acwing.com/problem/content/904/

package main

import (
	"fmt"
)

const N = 1010

var dp [N][N]int

func main() {
	var n, m int
	var s1, s2 string
	fmt.Scan(&n, &s1)
	fmt.Scan(&m, &s2)

	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= m; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 三者取最小
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + 1
			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+convertBool2Int(s1[i-1] != s2[j-1]))
		}
	}
	fmt.Println(dp[n][m])
}

func convertBool2Int(b bool) int {
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
