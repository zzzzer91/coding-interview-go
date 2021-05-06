// https://leetcode-cn.com/problems/longest-increasing-subsequence/

package main

func main() {

}

// 状态定义。dp[i] 表示以 nums[i] 结尾的最长上升子序列的长度。
// 即：在 [0, ..., i] 的范围内，选择以数字 nums[i] 结尾可以获得的最长上升子序列的长度。
func lengthOfLIS(nums []int) int {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
