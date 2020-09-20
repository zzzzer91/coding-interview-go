// https://leetcode-cn.com/problems/subsets-ii/

package main

import "sort"

func main() {

}

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var comb []int
	sort.Ints(nums)
	var dfs func(u int)
	dfs = func(u int) {
		temp := make([]int, len(comb))
		copy(temp, comb)
		res = append(res, temp)
		for i := u; i < len(nums); i++ {
			if i > u && nums[i] == nums[i-1] {
				continue
			}
			comb = append(comb, nums[i])
			dfs(i + 1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0)
	return res
}
