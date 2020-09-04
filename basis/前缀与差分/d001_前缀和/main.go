// https://www.acwing.com/problem/content/797/

package main

import (
	"fmt"
)

const N = 1e5 + 10

var a [N]int

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		a[i] += a[i-1]
	}
	var l, r int
	for i := 1; i <= m; i++ {
		fmt.Scan(&l, &r)
		fmt.Println(a[r] - a[l-1])
	}
}
