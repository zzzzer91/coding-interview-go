package main

import "fmt"

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10}
	selectionSort(a)
	fmt.Println(a)
}

func selectionSort(a []int) {
	for i, limit := 0, len(a)-1; i < limit; i++ {
		minIndex := i
		// 选择一个最小值的索引，然后交换
		for j := i + 1; j <= limit; j++ {
			if a[minIndex] > a[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			a[minIndex], a[i] = a[i], a[minIndex]
		}
	}
}
