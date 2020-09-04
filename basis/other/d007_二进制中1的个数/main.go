// https://www.acwing.com/problem/content/803/

package main

import (
	"fmt"
)

const N = 1e5 + 10

var a [N]int

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		c := binaryCount(a[i])
		if i == n-1 {
			fmt.Println(c)
		} else {
			fmt.Printf("%d ", c)
		}
	}
}

func binaryCount(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
