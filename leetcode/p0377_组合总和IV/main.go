// https://leetcode-cn.com/problems/combination-sum-iv/
// 注意，顺序不同的序列被视作不同的组合
// 应该算完全背包变种

package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	// dfs会超时
	// 使用dp数组，dp[i]代表组合数为i时使用nums中的数能组成的组合数的个数
	// dp[i]=dp[i-nums[0]]+dp[i-nums[1]]+dp[i=nums[2]]+...
	// 举个例子比如nums=[1,3,4],target=7;
	// dp[7]=dp[6]+dp[4]+dp[3]
	// 其实就是说7的组合数可以由三部分组成，1和dp[6]，3和dp[4],4和dp[3]
	for j := 1; j <= target; j++ {
		for _, n := range nums {
			if j >= n {
				dp[j] += dp[j-n]
			}
		}
	}
	return dp[target]
}

// 暴力，超时
func combinationSum4_2(nums []int, target int) int {
	res := 0
	var dfs func(target int)
	dfs = func(target int) {
		if target == 0 {
			res++
			return
		}
		for _, n := range nums {
			if n > target {
				continue
			}
			dfs(target - n)
		}
	}
	dfs(target)
	return res
}
