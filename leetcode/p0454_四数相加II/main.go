// https://leetcode-cn.com/problems/4sum-ii/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int, len(C)*len(D))
	for _, v1 := range C {
		for _, v2 := range D {
			m[v1+v2]++
		}
	}
	res := 0
	for _, v1 := range A {
		for _, v2 := range B {
			res += m[-v1-v2]
		}
	}
	return res
}
