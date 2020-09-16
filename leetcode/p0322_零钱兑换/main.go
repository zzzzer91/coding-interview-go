package main

import "fmt"

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ { // 注意 dp[0] 还是 0，作为初始值
		dp[i] = amount + 1 // 设定一个不可能达到的值作为最大值
	}
	for i := range coins {
		for j := coins[i]; j <= amount; j += 1 {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == amount+1 { // 没有到达，即凑不出组合，返回-1
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
