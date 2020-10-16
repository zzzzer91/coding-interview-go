// https://leetcode-cn.com/problems/merge-intervals/

package main

import "sort"

func main() {

}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		cur, pre := intervals[i], res[len(res)-1]
		if cur[0] <= pre[1] {
			res[len(res)-1][1] = max(pre[1], cur[1])
		} else {
			res = append(res, cur)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
