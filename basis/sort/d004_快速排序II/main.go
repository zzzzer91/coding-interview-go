// 简化快速排序代码，使其更容易记忆

package main

import (
	"fmt"
)

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10, 0}
	quickSort(a)
	fmt.Println(a)
}

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}

	l, r := 0, len(a)-1
	x := a[l+r>>1]
	for l < r {
		for a[l] < x {
			l++
		}
		for a[r] > x {
			r--
		}
		if l < r {
			a[l], a[r] = a[r], a[l]
		}
	}

	quickSort(a[:l])
	quickSort(a[l+1:])
}
