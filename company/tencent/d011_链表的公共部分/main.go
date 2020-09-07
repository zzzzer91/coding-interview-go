// 两个链表找交集。 AC

package main

import "fmt"

const N = 1e6 + 10

var a, b [N]int

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	var m int
	fmt.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
	}
	p1, p2 := 0, 0
	for p1 < n && p2 < m {
		if a[p1] > b[p2] {
			p1++
		} else if a[p1] < b[p2] {
			p2++
		} else {
			fmt.Printf("%d ", a[p1])
			p1++
			p2++
		}
	}
}
