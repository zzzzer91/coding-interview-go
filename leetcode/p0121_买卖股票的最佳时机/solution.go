package p0121_买卖股票的最佳时机

func maxProfit(prices []int) int {
	profit := 0
	last := 0
	for i := 0; i < len(prices)-1; i++ {
		last = max(0, last+prices[i+1]-prices[i])
		profit = max(profit, last)
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
