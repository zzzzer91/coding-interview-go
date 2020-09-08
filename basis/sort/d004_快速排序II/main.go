// 简化快速排序代码，使其更容易记忆
// todo 有问题

package main

import (
	"fmt"
)

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10, 0}
	// a := []int{2, 1}
	quickSort(a)
	fmt.Println(a)
}

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}

	l, r := 0, len(a)-1
	m := l + r>>1
	x := a[m]
	for l < r {
		for a[l] < x {
			l++
		}
		for a[r] > x {
			r--
		}
		if l < r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
	}

	quickSort(a[:m])
	quickSort(a[m+1:])
}
