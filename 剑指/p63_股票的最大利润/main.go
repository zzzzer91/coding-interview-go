// https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/

package main

import "math"

func main() {

}

// 只需要遍历价格数组一遍，记录历史最低点，
// 然后在每一天考虑这么一个问题：
// 如果我是在历史最低点买进的，那么我今天卖出能赚多少钱？
// 当考虑完所有天数之时，我们就得到了最好的答案。
func maxProfit(prices []int) int {
	res := 0
	minPrice := math.MaxInt64
	for _, p := range prices {
		minPrice = min(minPrice, p) // 找到有史以来的最低价格
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
