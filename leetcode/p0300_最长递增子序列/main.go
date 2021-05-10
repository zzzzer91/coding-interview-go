// https://leetcode-cn.com/problems/longest-increasing-subsequence/
// 题解：https://leetcode-cn.com/problems/longest-increasing-subsequence/

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

// 递归加记忆化
func lengthOfLIS(nums []int) int {
	m := make([]int, len(nums))
	var f func(i int) int
	// f(i) 代表以 nums[i] 开头的最长递增子序列
	f = func(i int) int {
		if v := m[i]; v > 0 {
			return v
		}
		res := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				res = max(res, f(j)+1)
			}
		}
		m[i] = res
		return res
	}
	res := 1
	for i := range nums {
		res = max(res, f(i))
	}
	return res
}

// 状态定义。dp[i] 表示以 nums[i] 结尾的最长上升子序列的长度。
// 即：在 [0, ..., i] 的范围内，选择以数字 nums[i] 结尾可以获得的最长上升子序列的长度。
func lengthOfLIS2(nums []int) int {
	res := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// 在遍历的时候同时找 dp 数组的最大值
		res = max(res, dp[i])
	}
	return res
}

// 状态定义。dp[i] 表示以 nums[i] 开头的最长上升子序列的长度。
func lengthOfLIS3(nums []int) int {
	res := 0
	dp := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
