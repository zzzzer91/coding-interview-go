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
			temp := dp[i-1][j-1]
			if s1[i-1] != s2[j-1] { // 替换
				temp += 1
			}
			// 三者取最小
			dp[i][j] = min3(dp[i][j-1]+1, dp[i-1][j]+1, temp)
		}
	}
	fmt.Println(dp[n][m])
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
