// https://leetcode-cn.com/problems/combination-sum/
// 顺序不同视为同一种组合，数组中每个元素都可以在同一组合中使用无数次

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
		if target == 0 {
			temp := make([]int, len(comb))
			copy(temp, comb)
			res = append(res, temp)
			return
		}
		for i := u; i < len(candidates); i++ {
			n := candidates[i]
			if n > target {
				continue
			}
			comb = append(comb, n)
			dfs(target-n, i) // 因为每个元素可以取无数次，所以 i 不变
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return res
}
