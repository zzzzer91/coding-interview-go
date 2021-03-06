// 快速排序
// 属于不稳定排序
// 平均时 O(nlog₂n)，最坏时 O(n²)，最好时 O(nlog₂n)
// 空 O(log₂n)

package main

import (
	"fmt"
)

func main() {
	a := []int{0, -1, 3, -5, 7, -9, 2, -4, 6, -8, 10, 0}
	a = []int{-1, -5, 2, 0, -5}
	quickSort(a)
	fmt.Println(a)
}

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	pivotIdx := partition(a)
	quickSort(a[:pivotIdx])
	quickSort(a[pivotIdx+1:])
}

func partition(a []int) int {
	l, r := 0, len(a)-1
	// 最左边数的作为枢纽，但效果不一定好
	pivot := a[l]
	// 最左边、中间、最右边的中间值作为枢纽
	// 同时中间值会被移动到最左边
	// pivot := median3(a)
	for l < r {
		// 从右往左，找到第一个小于 pivot 的值的索引
		for l < r && a[r] >= pivot {
			r--
		}
		// a[l] 的值目前暂存在 pivot 中，所以可以覆盖
		a[l] = a[r]
		// 从左往右，找到第一个大于 pivot 的值的索引
		for l < r && a[l] <= pivot {
			l++
		}
		a[r] = a[l]
	}
	// 此时 l == r，所以哪个都无所谓
	a[l] = pivot
	return l
}

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
