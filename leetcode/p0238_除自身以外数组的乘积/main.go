// https://leetcode-cn.com/problems/product-of-array-except-self/

package main

func main() {

}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = 1
	}
	lv, rv := 1, 1
	for i := range nums {
		res[i] *= lv
		lv *= nums[i]
		res[n-1-i] *= rv
		rv *= nums[n-1-i]
	}
	return res
}
