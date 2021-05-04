// https://leetcode-cn.com/problems/maximum-subarray/

package main

func main() {

}

// 动态规划
func maxSubArray(nums []int) int {
	res := nums[0]
	sum := nums[0]
	for _, n := range nums[1:] {
		sum = max(sum, 0) + n
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
