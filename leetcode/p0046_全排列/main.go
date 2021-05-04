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

// 更好的方法
// 单这样生成的全排列并不是按字典序存储在答案数组中的，
// 如果题目要求按字典序输出，那么还是用上面一种方法。
func permute2(nums []int) [][]int {
	var res [][]int
	var dfs func(u int)
	dfs = func(u int) {
		if u == len(nums)-1 {
			temp := make([]int, len(nums))
			copy(temp, nums)
			res = append(res, temp)
		}
		for i := u; i < len(nums); i++ {
			nums[i], nums[u] = nums[u], nums[i]
			dfs(u + 1)
			nums[i], nums[u] = nums[u], nums[i]
		}
	}
	dfs(0)
	return res
}
