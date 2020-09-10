// https://www.acwing.com/problem/content/844/

package main

import (
	"fmt"
)

const N = 10

var n int
var a [N]int
var st [N]bool

func main() {
	fmt.Scan(&n)
	dfs(0)
}

func dfs(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			if i == n-1 {
				fmt.Println(a[i])
			} else {
				fmt.Printf("%d ", a[i])
			}
		}
		return
	}
	for i := 1; i <= n; i++ { // 找没被用过的数字
		if st[i] {
			continue
		}
		a[u] = i
		st[i] = true // 代表这个数字已被使用
		dfs(u + 1)
		st[i] = false // 这个数字在这层已被使用完毕，释放掉，在另一层用
	}
}
