// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

package main

import "fmt"

func main() {
	// fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
	fmt.Println(search([]int{3, 5, 1}, 3))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] == target {
			return m
		}
		if nums[m] > nums[r] {
			if target < nums[m] && target > nums[r] { // [3, 5, 1]
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[m] && target <= nums[r] { // [5, 1, 3]
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}
