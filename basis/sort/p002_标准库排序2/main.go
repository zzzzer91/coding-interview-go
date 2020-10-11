package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "bc", "def"}
	// 自定义排序规则
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})
	fmt.Println(strs)
}
