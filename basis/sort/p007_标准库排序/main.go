package main

import (
	"fmt"
	"sort"
)

type Data [][]int

func (d Data) Len() int {
	return len(d)
}

func (d Data) Less(i, j int) bool {
	return d[i][0] < d[j][0] // < 代表升序，> 代表降序
}

func (d Data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	// 按照左端点排序
	a := [][]int{{3, 13}, {1, 11}, {2, 12}}
	sort.Sort(Data(a))
	fmt.Println(a)
}
