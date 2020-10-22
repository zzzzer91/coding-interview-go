// https://leetcode-cn.com/problems/partition-equal-subset-sum/submissions/

package main

func main() {

}

// 动态规划
func canPartition(nums []int) bool {
	sum := 0
	maxEle := -1
	for _, v := range nums {
		sum += v
		if maxEle < v {
			maxEle = v
		}
	}
	if (sum & 1) == 1 {
		return false
	}
	sum /= 2
	if maxEle > sum {
		return false
	}

	// 背包问题
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := range nums {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[sum]
}

// 超时
func canPartition_wrong(nums []int) bool {
	sum := 0
	maxEle := -1
	for _, v := range nums {
		sum += v
		if maxEle < v {
			maxEle = v
		}
	}
	if (sum & 1) == 1 {
		return false
	}
	sum /= 2
	if maxEle > sum {
		return false
	}

	var dfs func(u, target int) bool
	dfs = func(u, target int) bool {
		if target == 0 {
			return true
		}
		for i := u; i < len(nums); i++ {
			if nums[i] > target {
				continue
			}
			if dfs(i+1, target-nums[i]) {
				return true
			}
		}
		return false
	}
	return dfs(0, sum)
}
