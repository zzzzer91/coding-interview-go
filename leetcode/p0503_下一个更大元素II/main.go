// https://leetcode-cn.com/problems/next-greater-element-ii/

package main

func main() {

}

func nextGreaterElements(nums []int) []int {
	numsLen := len(nums)
	res := make([]int, numsLen)
	for i := range res {
		res[i] = -1
	}
	var stack []int
	for i := 0; i < numsLen*2-1; i++ {
		for len(stack) > 0 && nums[i%numsLen] > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = nums[i%numsLen]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%numsLen)
	}
	return res
}
