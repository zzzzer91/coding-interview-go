package p0581_最短无序连续子数组

import "sort"

// 先将数组排序，然后对比前后不同，需要辅助数组
func findUnsortedSubarray(nums []int) int {
	numsTemp := make([]int, len(nums))
	copy(numsTemp, nums)
	sort.Ints(numsTemp)

	lo, hi := 0, len(nums)-1
	// 注意是 <=，否则两个数组完全相同时，return会有问题
	for lo <= hi {
		if numsTemp[lo] != nums[lo] {
			break
		}
		lo++
	}
	for lo < hi {
		if numsTemp[hi] != nums[hi] {
			break
		}
		hi--
	}
	return hi - lo + 1
}

// 优化掉辅助数组
func findUnsortedSubarray2(nums []int) int {
	lo, hi := -1, -2
	highest := len(nums) - 1
	maxN, minN := nums[0], nums[highest]
	for i := 1; i <= highest; i++ {
		if maxN > nums[i] {
			// 注意，动的是 hi
			hi = i
		} else {
			maxN = nums[i]
		}
		if minN < nums[highest-i] {
			// 注意，动的是 lo
			lo = highest - i
		} else {
			minN = nums[highest-i]
		}
	}
	return hi - lo + 1
}
