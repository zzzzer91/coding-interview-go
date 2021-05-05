// https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/

package main

import "sort"

func main() {

}

// 先将数组排序，然后对比前后不同，需要辅助数组
func findUnsortedSubarray(nums []int) int {
	numsTemp := make([]int, len(nums))
	copy(numsTemp, nums)
	sort.Ints(numsTemp)

	l, r := 0, len(nums)-1
	// 注意是 <=，否则两个数组完全相同时，return会有问题
	for l <= r && numsTemp[l] == nums[l] {
		l++
	}
	for l < r && numsTemp[r] == nums[r] {
		r--
	}
	return r - l + 1
}

// 优化掉辅助数组
func findUnsortedSubarray2(nums []int) int {
	lo, hi := -1, -2 // 让 len(nums) <= 1 也成立
	numsLen := len(nums) - 1
	maxN, minN := nums[0], nums[numsLen]
	for i := 1; i <= numsLen; i++ {
		if maxN > nums[i] {
			// 注意，动的是 hi
			hi = i
		} else {
			maxN = nums[i]
		}
		if minN < nums[numsLen-i] {
			// 注意，动的是 lo
			lo = numsLen - i
		} else {
			minN = nums[numsLen-i]
		}
	}
	return hi - lo + 1
}
