// https://leetcode-cn.com/problems/132-pattern/

package main

func main() {

}

// 暴力
func find132pattern(nums []int) bool {
	for i := range nums {
		prev := nums[i]
		for j := i + 1; j < len(nums); j++ {
			prev = max(prev, nums[j])
			if nums[i] < nums[j] && nums[j] < prev {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
