// 动态规划

package main

func main() {

}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}

	dp := make([]int, n+1)
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 其实是上面的更简化
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	one, two := 1, 2
	current := 0
	for i := 3; i <= n; i++ {
		current = one + two
		one = two
		two = current
	}
	return current
}
