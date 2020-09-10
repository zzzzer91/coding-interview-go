package main

func main() {

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
