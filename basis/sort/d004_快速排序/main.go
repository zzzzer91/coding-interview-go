package main

import (
	"fmt"
)

// median3 选出数组最左边、中间、最右边的中间值，
// 并将中间值移动到最左边
func median3(a []int) int {
	l, r := 0, len(a)-1
	m := l + (r-l)/2
	medianIndex := l
	if (a[m]-a[l])^(a[m]-a[r]) < 0 { // 否则说明 a[m] 最小或最大
		medianIndex = m
	} else if (a[r]-a[l])^(a[r]-a[m]) < 0 { // 否则说明 a[r] 最小或最大
		medianIndex = r
	}
	// 将中间值移动到最左边（必须），因为采用的快排算法使用最左边值为 pivot
	if medianIndex != l {
		a[medianIndex], a[l] = a[l], a[medianIndex]
	}
	return a[l]
}

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}

	lo, hi := 0, len(a)-1
	// 最左边数的作为枢纽，但效果不一定好
	// pivot := a[lo]
	// 最左边、中间、最右边的中间值作为枢纽
	// 同时中间值会被移动到最左边
	pivot := median3(a)
	for lo < hi {
		// 从右往左，找到第一个小于 pivot 的值的索引
		for lo < hi && a[hi] >= pivot {
			hi--
		}
		// a[lo] 的值目前暂存在 pivot 中，所以可以覆盖
		a[lo] = a[hi]
		// 从左往右，找到第一个大于 pivot 的值的索引
		for lo < hi && a[lo] <= pivot {
			lo++
		}
		a[hi] = a[lo]
	}
	// 此时 lo == hi，所以哪个都无所谓
	a[lo] = pivot

	quickSort(a[:lo])
	quickSort(a[lo+1:])
}

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10}
	quickSort(a)
	fmt.Println(a)
}
