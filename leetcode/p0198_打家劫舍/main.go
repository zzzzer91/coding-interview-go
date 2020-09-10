package main

func main() {

}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := nums[0], nums[1]
	money := max(a, b)
	for _, num := range nums[2:] {
		money = max(a+num, b)
		a = max(a, b)
		b = money
	}
	return money
}

// 代码优化
func rob2(nums []int) int {
	preMoney, curMoney := 0, 0
	for _, num := range nums {
		preMoney, curMoney = curMoney, max(preMoney+num, curMoney)
	}
	return curMoney
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
