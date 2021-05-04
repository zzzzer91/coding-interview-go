// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

package main

func main() {

}

func maxProfit(prices []int) int {
	res := 0
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit = max(0, profit) + prices[i] - prices[i-1]
		res = max(res, profit)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
