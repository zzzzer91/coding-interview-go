// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/

package main

import "sort"

func main() {

}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		a, b := points[i], points[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
	res := 0
	var cur []int
	for _, p := range points {
		if cur != nil && cur[1] >= p[0] {
			cur[0] = p[0]
			if cur[1] > p[1] {
				cur[1] = p[1]
			}
		} else {
			cur = p
			res++
		}
	}
	return res
}

// 更优，按右端升序排序
func findMinArrowShots2(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	res := 1
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			res++
		}
	}
	return res
}
