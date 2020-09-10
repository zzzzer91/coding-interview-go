// https://www.acwing.com/problem/content/description/679/

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	n = 1024 - n
	count := 0
	coins := []int{64, 16, 4, 1}
	for _, c := range coins {
		count += n / c
		n %= c
	}
	fmt.Println(count)
}
