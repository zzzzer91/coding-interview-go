// https://leetcode-cn.com/problems/combination-sum/
// 候选集合中没有重复元素，且每个元素允许选无数次。
// 不能出现重复集合（顺序不同，值相同视为同一集合）

package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
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
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i) // 因为每个元素可以取无数次，所以 i 不变
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return res
}
