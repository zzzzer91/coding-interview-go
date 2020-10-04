// https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/

package main

import "math"

func main() {

}

func twoSum(n int) []float64 {
	var dp [12][100]int       // dp[i][j] 前 i 个骰子点数为 j 的情况下的总数
	for i := 1; i <= 6; i++ { // 先初始化一个骰子的情况
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ { // 从 2 个骰子开始迭代
		for j := i; j <= 6*n; j++ {
			for k := 1; k <= 6; k++ {
				if j-k >= 1 {
					dp[i][j] += dp[i-1][j-k]
				}
			}
		}
	}

	all := math.Pow(6, float64(n))
	var res []float64
	for i := n; i <= 6*n; i++ {
		if dp[n][i] != 0 {
			res = append(res, float64(dp[n][i])/all)
		}
	}
	return res
}

// 二维压缩成一维
func twoSum2(n int) []float64 {
	var dp [100]int           // dp[i][j] 前 i 个骰子点数为 j 的情况下的总数
	for i := 1; i <= 6; i++ { // 先初始化一个骰子的情况
		dp[i] = 1
	}
	for i := 2; i <= n; i++ { // 从 2 个骰子开始迭代
		for j := 6 * i; j >= i; j-- {
			dp[j] = 0
			for k := 1; k <= 6 && j-k >= i-1; k++ {
				dp[j] += dp[j-k]
			}
		}
	}

	all := math.Pow(6, float64(n))
	var res []float64
	for i := n; i <= 6*n; i++ {
		if dp[i] != 0 {
			res = append(res, float64(dp[i])/all)
		}
	}
	return res
}
