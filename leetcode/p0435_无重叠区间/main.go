// https://leetcode-cn.com/problems/non-overlapping-intervals/
// 贪心
// 题解：https://leetcode-cn.com/problems/non-overlapping-intervals/solution/tan-xin-suan-fa-zhi-qu-jian-diao-du-wen-ti-by-labu/

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
	return d[i][1] < d[j][1] // < 代表升序，> 代表降序
}

func (d Data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Sort(Data(intervals)) // 根据右端点排序
	end := intervals[0][1]
	count := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}
	return len(intervals) - count
}
