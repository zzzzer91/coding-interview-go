// https://leetcode-cn.com/problems/house-robber-ii/

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
	return max(f(nums[:len(nums)-1]), f(nums[1:]))
}

func f(nums []int) int {
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
