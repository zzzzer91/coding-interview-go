// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

package main

func main() {

}

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := nums[0]
	for _, n := range nums[1:] {
		if sum < 0 {
			sum = n
		} else {
			sum += n
		}
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
