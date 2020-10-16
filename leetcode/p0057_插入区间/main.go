// https://leetcode-cn.com/problems/insert-interval/

package main

func main() {

}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{newInterval}
	if len(intervals) == 0 {
		return res
	}
	for i := range intervals {
		a, b := intervals[i], res[len(res)-1]
		if a[0] <= b[1] {
			if a[1] >= b[0] {
				b[0], b[1] = min(a[0], b[0]), max(a[1], b[1])
			} else {
				res[len(res)-1] = a
				res = append(res, b)
			}
		} else {
			res = append(res, a)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
