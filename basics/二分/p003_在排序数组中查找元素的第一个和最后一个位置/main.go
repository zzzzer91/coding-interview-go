// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

func main() {

}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l := findFirstPos(nums, target)
	if l == -1 {
		return []int{-1, -1}
	}
	r := findLastPos(nums, target)
	return []int{l, r}
}

func findFirstPos(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}

func findLastPos(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l+1)>>1
		if nums[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}
