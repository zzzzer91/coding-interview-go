// https://leetcode-cn.com/problems/sort-array-by-parity-ii/

package main

func main() {

}

func sortArrayByParityII(A []int) []int {
	l, r := 0, len(A)-1
	for l < len(A) && r >= 0 {
		for l < len(A) && A[l]&1 == 0 {
			l += 2
		}
		for r >= 0 && A[r]&1 != 0 {
			r -= 2
		}
		if l < len(A) && r >= 0 {
			A[l], A[r] = A[r], A[l]
		}
	}
	return A
}
