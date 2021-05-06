// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/

package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}
