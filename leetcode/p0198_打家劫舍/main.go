// https://leetcode-cn.com/problems/house-robber/

package main

func main() {

}

// 代码优化
func rob(nums []int) int {
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
