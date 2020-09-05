// https://www.acwing.com/problem/content/551/

package main

import (
	"fmt"
	"sort"
)

var t int
var n, p int
var ss = make([]int, 1e5+10)
var sums = make([]int, 1e5+10)

func main() {
	fmt.Scanf("%d", &t)
	for i := 1; i <= t; i++ {
		fmt.Scanf("%d %d", &n, &p)
		for i := 1; i <= n; i++ {
			fmt.Scanf("%d", &ss[i])
		}
		sort.Ints(ss[1 : n+1])
		for i := 1; i <= n; i++ {
			sums[i] = sums[i-1] + ss[i]
		}
		res := int(1e9)
		// 滑动窗口思想，窗口大小p，每次前进1
		for i := p; i <= n; i++ {
			res = min(res, ss[i]*p-(sums[i]-sums[i-p]))
		}
		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
