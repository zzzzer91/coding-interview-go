// 给定长度为n的数组，依次删除每个数，输出剩下的数组成的数组的中位数。AC

package main

import (
	"fmt"
	"sort"
)

const N = 200010

var a = make([]int, N)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	at := make([]int, n)
	copy(at, a[:n])
	sort.Ints(at)
	m := n / 2
	for i := 0; i < n; i++ {
		if a[i] < at[m] {
			fmt.Printf("%d\n", at[m])
		} else {
			fmt.Printf("%d\n", at[m-1])
		}
	}
}
