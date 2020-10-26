// https://www.acwing.com/problem/content/156/

package main

import (
	"fmt"
)

const N = 1000010

var n, k int
var a, q [N]int
var ff, rr int

func main() {
	fmt.Scanf("%d%d", &n, &k)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	ff, rr = 0, -1 // 注意这里初始化
	for i := 0; i < n; i++ {
		if ff <= rr && i-q[ff] > k-1 {
			ff++
		}
		for ff <= rr && a[i] <= a[q[rr]] {
			rr--
		}
		rr++
		q[rr] = i
		if i >= k-1 {
			fmt.Printf("%d ", a[q[ff]])
		}
	}
	fmt.Println()

	ff, rr = 0, -1
	for i := 0; i < n; i++ {
		if ff <= rr && i-q[ff] > k-1 {
			ff++
		}
		for ff <= rr && a[i] >= a[q[rr]] {
			rr--
		}
		rr++
		q[rr] = i
		if i >= k-1 {
			fmt.Printf("%d ", a[q[ff]])
		}
	}
	fmt.Println()
}
