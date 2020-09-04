package p0045_跳跃游戏II

func jump(nums []int) int {
	res := 0
	end := 0
	maxPos := 0
	// len(nums)-1，防止正好跳到末尾的情况
	for i, num := range nums[:len(nums)-1] {
		//找能跳的最远的
		maxPos = max(maxPos, num+i)
		if i == end { // 遇到边界，就更新边界，并且步数加一
			end = maxPos
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
