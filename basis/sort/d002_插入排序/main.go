package main

import "fmt"

// 可视化见 https://visualgo.net/en/sorting
func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		j := i
		for ; j > 0 && temp < a[j-1]; j-- {
			a[j] = a[j-1]
		}
		a[j] = temp
	}
}

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10}
	insertionSort(a)
	fmt.Println(a)
}
