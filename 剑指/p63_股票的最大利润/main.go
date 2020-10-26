// https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/

package main

import "math"

func main() {

}

func maxProfit(prices []int) int {
	res := 0
	minPrice := math.MaxInt64
	for _, p := range prices {
		minPrice = min(minPrice, p)
		res = max(res, p-minPrice)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit2(prices []int) int {
	res := 0
	sum := 0
	for i := 1; i < len(prices); i++ {
		if sum < 0 {
			sum = prices[i] - prices[i-1]
		} else {
			sum += prices[i] - prices[i-1]
		}
		res = max(res, sum)
	}
	return res
}
