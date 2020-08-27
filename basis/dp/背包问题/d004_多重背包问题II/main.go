package main

import (
	"fmt"
)

const N = 1010

var vs, ws [N * 11]int
var ss [N]int
var dp [N * 2]int

func main() {
	var n, v int
	fmt.Scan(&n, &v)
	count := n + 1
	for i := 1; i <= n; i++ {
		fmt.Scan(&vs[i], &ws[i], &ss[i])
		t1, t2 := vs[i], ws[i]
		sum := 1
		for j := 2; sum+j <= ss[i]; j *= 2 { // 二进制拆分，转换为 01背包
			vs[count] = t1 * j
			ws[count] = t2 * j
			sum += j
			count++
		}
		if sum < ss[i] {
			vs[count] = t1 * (ss[i] - sum)
			ws[count] = t2 * (ss[i] - sum)
			count++
		}
	}

	for i := 0; i < count; i++ {
		for j := v; j >= vs[i]; j-- {
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
