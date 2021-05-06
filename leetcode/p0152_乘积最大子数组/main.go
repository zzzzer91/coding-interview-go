// https://leetcode-cn.com/problems/maximum-product-subarray/

package main

import "math"

func main() {

}

func maxProduct(nums []int) int {
	res := math.MinInt64
	iMax, iMin := 1, 1
	for _, n := range nums {
		if n < 0 { // 如果 a 小于 0，那么原本 min 会变成max
			iMax, iMin = iMin, iMax
		}
		iMax = max(iMax*n, n)
		iMin = min(iMin*n, n)
		res = max(res, iMax)
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
