// https://leetcode-cn.com/problems/interval-list-intersections/

package main

func main() {

}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	var res [][]int
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		a, b := A[i], B[j]
		if a[1] < b[1] {
			i++
		} else {
			a, b = b, a
			j++
		}
		if a[1] >= b[0] {
			res = append(res, []int{max(a[0], b[0]), a[1]})
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
