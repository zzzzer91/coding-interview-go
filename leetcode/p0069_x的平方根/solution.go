package p0069_x的平方根

// 二分查找法
func mySqrt(x int) int {
	// 这里必须检查
	if x == 1 {
		return 1
	}
	// low 从 0 开始，因为 x 可能 0
	low, high := 0, x
	for {
		mid := low + (high-low)/2
		if mid == low {
			return mid
		}
		square := mid * mid
		if square > x {
			high = mid
		} else if square < x {
			low = mid
		} else {
			return mid
		}
	}
}
