// 冒泡排序
// 属于稳定排序
// 平均时 O(n²)，最坏时 O(n²)，最好时 O(n)
// 空 O(1)

package main

import "fmt"

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10}
	bubbleSort(a)
	fmt.Println(a)
}

func bubbleSort(a []int) {
	flag := false
	for i := len(a) - 1; i > 0; i-- {
		flag = false
		// 每轮将剩余数组中最大的数移到最后
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		// flag 为 false，说明上面一趟没有发生排序，即所有元素已经排好序
		if !flag {
			break
		}
	}
}
