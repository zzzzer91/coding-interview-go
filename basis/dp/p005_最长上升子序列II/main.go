// https://www.acwing.com/problem/content/898/
// 状态表示：f[i]表示从第一个数字开始算， 以w[i]结尾的最大的上升序列。
// (以w[i]结尾的所有上升序列中属性为最大值的那一个)
// 与 I 比，数据范围更大，所以我们需要一种更效率的算法

package main

import (
	"fmt"
)

const N = 1001

var a [N]int
var dp [N]int

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
