package main

import "sort"

func main() {

}

// 先根据前一个值排序
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Sort(data(intervals))
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}

type data [][]int

func (d data) Len() int {
	return len(d)
}

func (d data) Less(i, j int) bool {
	return d[i][0] < d[j][0]
}

func (d data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
