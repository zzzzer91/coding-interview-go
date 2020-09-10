// https://leetcode-cn.com/problems/combination-sum-ii/
// 相比较 p0039，每个元素只能用一次

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2_notRight([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 先要排序
	var res [][]int
	var comb []int
	var dfs func(target, u int)
	dfs = func(target, u int) {
		if target == 0 {
			temp := make([]int, len(comb))
			copy(temp, comb)
			res = append(res, temp)
			return
		}
		for i := u; i < len(candidates); i++ {
			// 剪枝
			if candidates[i] > target {
				return
			}
			// 当前分支和之前分支相同，剪掉
			// 因为当前元素和之前元素值相同的话，继续往下构建的分支一定是一样的
			if i > u && candidates[i] == candidates[i-1] {
				continue
			}
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return res
}

// 错误示范：这种方法并不能保证每个元素只取一次
func combinationSum2_notRight(candidates []int, target int) [][]int {
	var res [][]int
	var comb []int
	var dfs func(target, u int)
	dfs = func(target, u int) {
		if target < 0 {
			return
		}
		if target == 0 {
			temp := make([]int, len(comb))
			copy(temp, comb)
			res = append(res, temp)
			return
		}
		for i := u; i < len(candidates); i++ {
			// 剪枝
			if candidates[i] > target {
				continue
			}
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return res
}
