// https://www.acwing.com/problem/content/description/758/

package main

import (
	"fmt"
)

const N = 110

var n, m int
var a [N][N]int

func main() {
	fmt.Scanf("%d %d", &n, &m)
	ll, rr, tt, bb := 0, m-1, 0, n-1
	k := 1
	for ll <= rr && tt <= bb {
		for i := ll; i <= rr; i++ {
			a[tt][i] = k
			k++
		}
		for i := tt + 1; i <= bb; i++ {
			a[i][rr] = k
			k++
		}
		for i := rr - 1; i >= ll; i-- {
			a[bb][i] = k
			k++
		}
		for i := bb - 1; i > tt; i-- {
			a[i][ll] = k
			k++
		}
		ll++
		rr--
		tt++
		bb--
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Printf("\n")
	}
}
