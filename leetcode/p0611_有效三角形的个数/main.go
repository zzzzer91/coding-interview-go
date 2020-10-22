// https://leetcode-cn.com/problems/valid-triangle-number/

package main

import "sort"

func main() {

}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := len(nums) - 1; i >= 2; i-- {
		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				res += r - l // l 右边的数不需要再算了，因为升序排序过了，必然大于
				r--
			} else {
				l++
			}
		}
	}
	return res
}
