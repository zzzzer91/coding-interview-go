// https://leetcode-cn.com/problems/squares-of-a-sorted-array/

package main

func main() {

}

func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	l, r := 0, len(A)-1
	pos := r
	for l <= r {
		if abs(A[r]) > abs(A[l]) {
			res[pos] = A[r] * A[r]
			pos--
			r--
		} else {
			res[pos] = A[l] * A[l]
			pos--
			l++
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
