// https://leetcode-cn.com/problems/permutations/
// 允许有重复序列

package main

func main() {

}

func permute(nums []int) [][]int {
	var res [][]int
	var comb []int
	st := make([]bool, len(nums))
	var dfs func(u int)
	dfs = func(u int) {
		if u == len(nums) {
			temp := make([]int, len(comb))
			copy(temp, comb)
			res = append(res, temp)
			return
		}
		for i, v := range nums {
			if st[i] {
				continue
			}
			comb = append(comb, v)
			st[i] = true
			dfs(u + 1)
			comb = comb[:len(comb)-1]
			st[i] = false
		}
	}
	dfs(0)
	return res
}
