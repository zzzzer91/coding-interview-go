// https://www.acwing.com/problem/content/9/

package main

import (
	"fmt"
)

const N = 101

var ss [N]int
var vs, ws [N][N]int
var dp [N][N]int

func main() {
	var n, v int
	fmt.Scan(&n, &v)
	for i := 1; i <= n; i++ {
		fmt.Scan(&ss[i])
		for j, length := 0, ss[i]; j < length; j++ {
			fmt.Scan(&vs[i][j], &ws[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= v; j++ {
			dp[i][j] = dp[i-1][j] // 要先赋值，这是不选的情况，下面max中比较要用到
			for k, length := 0, ss[i]; k < length; k++ {
				if j >= vs[i][k] {
					dp[i][j] = max(dp[i][j], dp[i-1][j-vs[i][k]]+ws[i][k])
				}
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
