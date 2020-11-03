// https://leetcode-cn.com/problems/valid-mountain-array/

package main

func main() {

}

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	pos := 1
	for pos < len(A) {
		if A[pos] < A[pos-1] {
			if pos == 1 {
				return false
			}
			break
		} else if A[pos] == A[pos-1] {
			return false
		}
		pos++
	}
	if pos == len(A) {
		return false
	}
	for pos < len(A) {
		if A[pos] >= A[pos-1] {
			return false
		}
		pos++
	}
	return true
}
