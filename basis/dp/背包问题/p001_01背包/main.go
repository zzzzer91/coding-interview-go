// https://www.acwing.com/problem/content/2/

package main

import "fmt"

const N = 1010

var vs, ws [N]int
var dp [N][N]int // dp含义，前 i 个物品，且总体积不超过 j 的选法集合

func main() {
	var n, v int
	fmt.Scan(&n, &v)
	for i := 1; i <= n; i++ {
		fmt.Scan(&vs[i], &ws[i])
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= v; j++ { // j 代表当前体积
			if j >= vs[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-vs[i]]+ws[i])
			} else {
				dp[i][j] = dp[i-1][j] // 物品放不下
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
