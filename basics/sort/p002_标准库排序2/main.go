package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][]int{{2, 5}, {2, 3}, {1, 6}, {4, 10}}
	// 自定义排序规则
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] < a[j][1]
		}
		return a[i][0] < a[j][0]
	})
	fmt.Println(a)
}
