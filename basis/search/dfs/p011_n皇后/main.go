// https://www.acwing.com/problem/content/845/
// 时间复杂度 n!

package main

import (
	"fmt"
)

const N = 10

var n int
var a [N][N]byte
var col [N]bool
var dg, udg [N * 2]bool // 要大于两倍数据范围，因为下面对角线判断，为了防止负数，要加上 n

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = '.'
		}
	}
	dfs(0)
}

func dfs(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			// 低版本 go，这行代码有问题？
			// 只输出一次就退出了
			fmt.Printf("%s\n", a[i])
		}
		fmt.Println()
		return
	}
	for j := 0; j < n; j++ {
		// y = -x + b   =>  b = y + x 这样就能唯一表示一个对角线上的一点
		// y =  x + b   =>  b = y k - x 有可能为负，所以加上偏移量n，不过这样数组范围会变大
		if col[j] || dg[j+u] || udg[j-u+n] {
			continue
		}
		a[u][j] = 'Q'
		col[j], dg[u+j], udg[j-u+n] = true, true, true
		dfs(u + 1)
		a[u][j] = '.'
		col[j], dg[u+j], udg[j-u+n] = false, false, false
	}
}
