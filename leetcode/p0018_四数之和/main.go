// https://leetcode-cn.com/problems/4sum/

package main

import "sort"

func main() {

}

// dfs，效率低
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var comb []int
	var dfs func(u, target int)
	dfs = func(u, target int) {
		if len(comb) == 4 {
			if target == 0 {
				temp := make([]int, len(comb))
				copy(temp, comb)
				res = append(res, temp)
			}
			return
		}
		for i := u; i < len(nums); i++ {
			if i > u && nums[i] == nums[i-1] {
				continue
			}
			comb = append(comb, nums[i])
			dfs(i+1, target-nums[i])
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0, target)
	return res
}

// 双指针优化
func fourSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				if l > j+1 && nums[l] == nums[l-1] {
					l++
					continue
				}
				if r < len(nums)-1 && nums[r] == nums[r+1] {
					r--
					continue
				}
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				}
			}
		}
	}
	return res
}
