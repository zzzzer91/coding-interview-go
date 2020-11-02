// https://www.acwing.com/problem/content/description/2/
// dp 压缩成一维，并省略数组

package main

import "fmt"

const N = 1001

var dp [N]int

func main() {
	var n, v int
	fmt.Scan(&n, &v)
	var volume, weight int
	for i := 0; i < n; i++ {
		fmt.Scan(&volume, &weight)
		for j := v; j >= volume; j-- {
			dp[j] = max(dp[j], dp[j-volume]+weight)
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
