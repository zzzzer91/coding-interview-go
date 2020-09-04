package p0042_接雨水

// 使用双指针
func trap(height []int) int {
	ret := 0
	l, r := 0, len(height)-1
	maxL, maxR := 0, 0
	for l < r {
		lV, rV := height[l], height[r]
		if lV < rV {
			if lV >= maxL {
				maxL = lV
			} else {
				ret += maxL - lV
			}
			l++
		} else {
			if rV >= maxR {
				maxR = rV
			} else {
				ret += maxR - rV
			}
			r--
		}
	}
	return ret
}
