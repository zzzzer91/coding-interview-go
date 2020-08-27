// https://www.acwing.com/problem/content/description/3/

package main

import "fmt"

const N = 1001

var vs, ws [N]int
var dp [N]int

func main() {
	var n, v int
	fmt.Scan(&n, &v)
	for i := 0; i < n; i++ {
		fmt.Scan(&vs[i], &ws[i])
	}

	for i := 0; i < n; i++ {
		// 与 01背包 的唯一区别，从前往后
		for j := vs[i]; j <= v; j++ {
			dp[j] = max(dp[j], dp[j-vs[i]]+ws[i])
		}
	}
	fmt.Println(dp[v])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
