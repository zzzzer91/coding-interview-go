// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		temp := prices[i] - prices[i-1]
		if temp > 0 {
			profit += temp
		}
	}
	return profit
}
