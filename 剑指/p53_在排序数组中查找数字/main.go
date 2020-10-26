// https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

package main

import "fmt"

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	r := findLastIdx(nums, target)
	if r == -1 {
		return 0
	}
	l := findFirstIdx(nums, target)
	return r - l + 1
}

func findLastIdx(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l+1)>>1
		if target < nums[m] {
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

func findFirstIdx(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if target <= nums[m] {
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
