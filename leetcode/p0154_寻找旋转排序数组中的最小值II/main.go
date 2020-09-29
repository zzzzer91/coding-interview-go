// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

package main

func main() {

}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) >> 1
		if nums[m] < nums[r] {
			r = m // 考虑 [1, 2] 这种情况
		} else if nums[m] > nums[r] {
			l = m + 1 // 考虑 [2, 1] 这种情况
		} else {
			r--
		}
	}
	return nums[l]
}
