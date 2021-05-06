// https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findUnsortedSubarray2([]int{2, 6, 4, 8, 10, 9, 15}))
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
	highest := len(nums) - 1
	preMaxN, preMinN := nums[0], nums[highest]
	for i := 1; i <= highest; i++ { // 注意是 <=
		// 从左到右找需要调整的最大位置 i
		// 注意，动的是 hi
		if nums[i] < preMaxN {
			hi = i
		} else {
			preMaxN = nums[i]
		}
		// 从右到左找需要调整的最小位置 i
		// 注意，动的是 lo
		if nums[highest-i] > preMinN {
			lo = highest - i
		} else {
			preMinN = nums[highest-i]
		}
	}
	return hi - lo + 1
}
