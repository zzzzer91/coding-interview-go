package main

import "fmt"

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10}
	insertionSort(a)
	fmt.Println(a)
}

// 可视化见 https://visualgo.net/en/sorting
func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		temp := a[i] // 暂存该值
		pos := i
		for pos > 0 && temp < a[pos-1] { // 将该值之前比该值大的元素往后诺
			a[pos] = a[pos-1]
			pos--
		}
		a[pos] = temp
	}
}
