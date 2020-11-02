// 有 n 个部队，部队人数小于等于另一个部队人数三分之一，可以藏在该部队后面不被发现，不可嵌套
// 问：最小可以只看到几个部队

package main

import (
	"fmt"
	"sort"
)

const N = 50010

var n int
var a [N]int

func main() {
	n = 5
	a[0], a[1], a[2], a[3], a[4] = 2, 6, 7, 7, 10
	fmt.Println(f())
}

func f() int {
	res := n
	sort.Ints(a[:n])
	i, j := (n-1)/2, n-1
	for i >= 0 && j >= (n-1)/2 {
		if a[i]*3 <= a[j] {
			res--
			i--
			j--
		} else {
			i--
		}
	}
	return res
}
