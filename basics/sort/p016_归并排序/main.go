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
	merge(a, m)
}

// merge 将数组被 mid 分为左右的两部分合并有序
func merge(a []int, m int) {
	temp := make([]int, len(a))
	copy(temp, a)
	pos := 0
	i, j := 0, m+1 // 应该用 m+1 而不是 m，考虑长度只有 2 的情况
	for i <= m && j < len(a) {
		if temp[i] < temp[j] {
			a[pos] = temp[i]
			i++
		} else {
			a[pos] = temp[j]
			j++
		}
		pos++
	}
	// 处理 temp 中剩下的元素
	if i <= m {
		copy(a[pos:], temp[i:m+1])
	} else {
		copy(a[pos:], temp[j:])
	}
}
