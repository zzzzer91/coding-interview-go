package main

import (
	"fmt"
)

const N = 110

var vs, ws, ss [N]int
var dp [N][N]int

func main() {
	var n, v int
	fmt.Scan(&n, &v)
	for i := 1; i <= n; i++ {
		fmt.Scan(&vs[i], &ws[i], &ss[i])
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= v; j++ {
			// dp[i][j] = dp[i-1][j] // 不选情况，可以合并到下面去，即 k=0
			for k := 0; k <= ss[i] && k*vs[i] <= j; k++ {
				dp[i][j] = max(dp[i][j], dp[i-1][j-k*vs[i]]+k*ws[i])
			}
		}
	}
	fmt.Println(dp[n][v])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
