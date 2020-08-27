// https://www.acwing.com/problem/content/900/

package main

import (
	"fmt"
)

const N = 501

var a [N][N]int
var dp [N]int

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = max(dp[j], dp[j+1]) + a[i][j]
		}
	}
	fmt.Println(dp[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
