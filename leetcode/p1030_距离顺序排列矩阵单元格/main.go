// https://leetcode-cn.com/problems/matrix-cells-in-distance-order/

package main

func main() {

}

// 桶排序
func allCellsDistOrder(R, C, r0, c0 int) [][]int {
	maxDist := max(r0, R-1-r0) + max(c0, C-1-c0)
	buckets := make([][][]int, maxDist+1)

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dist := abs(i-r0) + abs(j-c0)
			buckets[dist] = append(buckets[dist], []int{i, j})
		}
	}

	res := make([][]int, 0, R*C)
	for _, bucket := range buckets {
		res = append(res, bucket...)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
