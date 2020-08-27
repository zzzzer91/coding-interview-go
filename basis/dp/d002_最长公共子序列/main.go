// https://www.acwing.com/problem/content/899/
// 子序列不需要连续
// 题解：https://www.acwing.com/solution/content/8111

package main

import (
	"fmt"
)

const N = 1010

var dp [N][N]int

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var s1, s2 []byte
	fmt.Scan(&s1, &s2)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] { // 从0开始，所以减1
				dp[i][j] = dp[i-1][j-1] + 1 // dp从1开始
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	fmt.Println(dp[n][m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
