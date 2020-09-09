// 和上一题硬币问题很像
// 数组中每个数都可以取无限次，求和为 target 的方案数

package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // 2
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))    // 3
}

func combinationSum(candidates []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // 初始状态
	for i := range candidates {
		for j := candidates[i]; j <= target; j++ {
			dp[j] += dp[j-candidates[i]]
		}
	}
	return dp[target]
}
