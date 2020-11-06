// https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/

package main

import "sort"

func main() {

}

var bit = [1e4 + 1]int{}

func init() {
	for i := 1; i <= 1e4; i++ {
		bit[i] = bit[i>>1] + i&1
	}
}

func sortByBits(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := bit[x], bit[y]
		return cx < cy || cx == cy && x < y
	})
	return a
}
