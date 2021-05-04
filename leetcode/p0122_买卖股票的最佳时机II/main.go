// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			res += profit
		}
	}
	return res
}
