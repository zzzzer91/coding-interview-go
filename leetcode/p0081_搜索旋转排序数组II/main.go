// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/

package main

func main() {

}

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return true
		}
		if nums[m] < nums[r] { // 右半部分必然有序
			if target > nums[m] && target <= nums[r] { // 检查 target 是不是在右半部分
				l = m + 1
			} else {
				r = m - 1
			}
		} else if nums[m] > nums[r] { // 左半部分必然有序
			if target < nums[m] && target > nums[r] { // 检查 target 是不是在左半部分
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			r--
		}
	}
	return false
}
