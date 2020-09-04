package p0053_最大子序和

// 动态规划
func maxSubArray(nums []int) int {
	res := nums[0]
	sum := nums[0]
	for _, num := range nums[1:] {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
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
