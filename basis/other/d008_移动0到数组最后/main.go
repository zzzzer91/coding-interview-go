// 移动0到数组最后，并保持数组元素相对未知

package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 0, 3, 4}
	move(a)
}

func move(a []int) {
	pos := 0
	for i := range a {
		if a[i] != 0 {
			a[pos] = a[i]
			pos++
		}
	}
	for pos < len(a) {
		a[pos] = 0
		pos++
	}
	fmt.Println(a)
}
