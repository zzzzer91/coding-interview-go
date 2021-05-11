// 归并排序
// 属于稳定排序
// 平均时 O(nlog₂n)，最坏时 O(nlog₂n)，最好时 O(nlog₂n)
// 空 O(n)

package main

import "fmt"

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10}
	mergeSort(a)
	fmt.Println(a)
}

func mergeSort(a []int) {
	if len(a) <= 1 {
		return
	}
	m := (len(a) - 1) >> 1
	// 使左半部分有序
	mergeSort(a[:m+1]) // 应该用 m+1 而不是 m，考虑长度只有 2 的情况
	// 使右半部分有序
	mergeSort(a[m+1:])
	if a[m] <= a[m+1] { // 左半部分最大小于等于右半部分最小，即已经有序了，不需要再合并
		return
	}
	merge(a, m)
}

// merge 将数组被 mid 分为左右的两部分合并有序
func merge(a []int, m int) {
	temp := make([]int, len(a)) // temp 如果共用的话可以节约内存
	copy(temp, a)
	i, j, k := 0, m+1, 0 // 应该用 m+1 而不是 m，考虑长度只有 2 的情况
	for i <= m && j < len(a) {
		if temp[i] <= temp[j] { // 这里必须是小于等于，否则是不稳定排序
			a[k] = temp[i]
			i++
		} else {
			a[k] = temp[j]
			j++
		}
		k++
	}
	// 处理 temp 中剩下的元素
	if i <= m {
		copy(a[k:], temp[i:m+1])
	} else {
		copy(a[k:], temp[j:])
	}
}
