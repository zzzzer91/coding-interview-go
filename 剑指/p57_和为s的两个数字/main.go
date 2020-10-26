// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/

package main

func main() {

}

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return []int{nums[l], nums[r]}
		}
	}
	return []int{-1, -1}
}
