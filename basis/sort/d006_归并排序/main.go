package main

import "fmt"

// merge 将数组被 mid 分为左右的两部分合并有序
func merge(a []int, mid int) {
	temp := make([]int, len(a))
	copy(temp, a)
	pos := 0
	i, j := 0, mid+1
	for i <= mid && j < len(a) {
		if temp[i] < temp[j] {
			a[pos] = temp[i]
			i++
		} else {
			a[pos] = temp[j]
			j++
		}
		pos++
	}
	// 处理剩下的元素
	// 下面两个循环只可能执行一个
	for i <= mid {
		a[pos] = temp[i]
		i++
		pos++
	}
	for j < len(a) {
		a[pos] = temp[j]
		j++
		pos++
	}
}

func mergeSort(a []int) {
	if len(a) <= 1 {
		return
	}
	mid := (len(a) - 1) / 2
	mergeSort(a[:mid+1]) // 使左半部分有序
	mergeSort(a[mid+1:]) // 使右半部分有序
	merge(a, mid)
}

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10}
	mergeSort(a)
	fmt.Println(a)
}
