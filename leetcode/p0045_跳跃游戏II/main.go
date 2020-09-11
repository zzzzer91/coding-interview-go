// https://leetcode-cn.com/problems/jump-game-ii/
// 此题假设一定能跳到最后一个点，如果给的数据跳不到，结果会有问题

package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

// 贪心
// 比起 p0055 的贪心做法，需要额外变量记录边界，因为要输出次数
func jump(nums []int) int {
	res := 0
	end := 0
	maxPos := 0
	// 在遍历数组时，我们不访问最后一个元素，
	// 这是因为在访问最后一个元素之前，我们的边界一定大于等于最后一个位置，
	// 否则就无法跳到最后一个位置了。
	// 如果访问最后一个元素，在边界正好为最后一个位置的情况下，
	// 我们会增加一次「不必要的跳跃次数」，因此我们不必访问最后一个元素。
	for i, n := range nums[:len(nums)-1] {
		// 找直到 end 前，能跳的最远距离
		maxPos = max(maxPos, n+i)
		if i == end { // 遇到边界，就更新边界，并且最少需要步数加一
			end = maxPos // 更新边界为能跳到的最远距离
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
