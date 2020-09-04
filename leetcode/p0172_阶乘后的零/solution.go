package p0172_阶乘后的零

// count = n / 5 + n / 25 + n / 125 + ...
func trailingZeroes(n int) int {
	count := 0
	for i := 5; i <= n; i *= 5 {
		count += n / i
	}
	return count
}

// 递归法
func trailingZeroes2(n int) int {
	temp := n / 5
	if temp == 0 {
		return 0
	}
	return temp + trailingZeroes2(temp)
}
