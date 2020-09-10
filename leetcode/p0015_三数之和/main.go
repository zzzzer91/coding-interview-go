package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		// 去除重复对
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 去除重复对
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return res
}
