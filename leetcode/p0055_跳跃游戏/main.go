package main

func main() {

}

func canJump(nums []int) bool {
	// 直到 i-1 位置，能跳到的最远距离
	pre := nums[0]
	for _, num := range nums[1:] {
		// 跳不到下一格
		if pre < 1 {
			return false
		}
		// 更新当前能跳到的最远距离
		pre = max(num, pre-1)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
