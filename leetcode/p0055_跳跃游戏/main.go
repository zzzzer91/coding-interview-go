// https://leetcode-cn.com/problems/jump-game/

package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

// dp
func canJump(nums []int) bool {
	dp := make([]int, len(nums)) // 代表当前位置 i 能跳到的最远距离
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] == 0 { // 之前最大跳跃距离为 0，已经跳不到当前位置了
			return false
		}
		dp[i] = max(nums[i], dp[i-1]-1)
	}
	return true
}

// dp 压缩数组
func canJump2(nums []int) bool {
	// 直到 i-1 位置，能跳到的最远距离
	pre := nums[0]
	for _, num := range nums[1:] {
		// 跳不到下一格
		if pre == 0 {
			return false
		}
		// 更新当前能跳到的最远距离
		pre = max(num, pre-1)
	}
	return true
}

// 贪心
func canJump3(nums []int) bool {
	maxPos := 0 // 能跳到的最远下标
	for i, n := range nums {
		if maxPos < i { // 跳不到当前位置
			break
		}
		maxPos = max(maxPos, n+i)  // 更新能跳到的最远位置
		if maxPos >= len(nums)-1 { // 已经可以跳到最后一个位置了
			return true
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
