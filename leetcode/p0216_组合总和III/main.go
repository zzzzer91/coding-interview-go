// https://leetcode-cn.com/problems/combination-sum-iii/

package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
}

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var comb []int
	var dfs func(target, u int)
	dfs = func(target, u int) {
		if len(comb) == k {
			if target == 0 {
				temp := make([]int, len(comb))
				copy(temp, comb)
				res = append(res, temp)
			}
			return
		}
		for i := u; i <= 9; i++ {
			if i > target {
				break
			}
			comb = append(comb, i)
			dfs(target-i, i+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(n, 1)
	return res
}
