// 插入排序
// 属于稳定排序
// 平均时 O(n²)，最坏时 O(n²)，最好时 O(n)
// 空 O(1)

package main

import "fmt"

func main() {
	a := []int{0, -1, 3, -5, 7, -9, 2, -4, 6, -8, 10}
	insertionSort(a)
	fmt.Println(a)
}

// 可视化见 https://visualgo.net/en/sorting
func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		temp := a[i] // 暂存该值
		var pos int
		for pos = i; pos > 0 && temp < a[pos-1]; pos-- { // 将该值之前比该值大的元素往后移动
			a[pos] = a[pos-1]
		}
		a[pos] = temp
	}
}
