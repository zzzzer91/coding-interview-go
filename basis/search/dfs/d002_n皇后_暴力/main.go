// https://www.acwing.com/problem/content/845/

package main

import (
	"fmt"
)

const N = 10

var n int
var a [N][N]byte
var row, col [N]bool
var dg, udg [N * 2]bool // 要大于两倍数据范围，因为下面对角线判断，为了防止负数，要加上 n

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = '.'
		}
	}
	dfs(0, 0, 0)
}

// s 代表当前皇后数量
func dfs(x, y, s int) {
	if y == n {
		// 换行
		y = 0
		x++
	}
	if x == n { // 超过边界，结束
		if s == n {
			for i := 0; i < n; i++ {
				// 低版本 go，这行代码有问题？
				// 只输出一次就退出了
				fmt.Printf("%s\n", a[i])
			}
			fmt.Println()
		}
		return
	}

	// 两种情况都完全搜一遍

	// 不放皇后情况
	dfs(x, y+1, s)

	// 放皇后情况
	if !row[x] && !col[y] && !dg[x+y] && !udg[x-y+n] {
		a[x][y] = 'Q'
		row[x], col[y], dg[x+y], udg[x-y+n] = true, true, true, true
		dfs(x, y+1, s+1)
		a[x][y] = '.'
		row[x], col[y], dg[x+y], udg[x-y+n] = false, false, false, false
	}
}
