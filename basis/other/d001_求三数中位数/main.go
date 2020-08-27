package main

import (
	"fmt"
)

func median3(a, b, c int) int {
	median := a
	// 这里使用了异或，负数符号位为 1，正数符号位为 0，异或后还是负的
	if (b-a)^(b-c) < 0 { // 否则说明 b 最小或最大
		median = b
	} else if (c-a)^(c-b) < 0 { // 否则说明 c 最小或最大
		median = c
	}
	return median
}

func main() {
	fmt.Println(median3(1, 2, 3))
	fmt.Println(median3(1, 3, 2))
	fmt.Println(median3(2, 1, 3))
	fmt.Println(median3(2, 3, 1))
	fmt.Println(median3(3, 1, 2))
	fmt.Println(median3(3, 2, 1))

	fmt.Println(median3(3, 2, 3))

	fmt.Println(median3(-1, -2, -3))
}
