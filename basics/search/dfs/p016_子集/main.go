// https://leetcode-cn.com/problems/subsets/
// 集合中没有重复元素

package main

func main() {

}

func subsets(nums []int) [][]int {
	var res [][]int
	var comb []int
	var dfs func(u int)
	dfs = func(u int) {
		temp := make([]int, len(comb))
		copy(temp, comb)
		res = append(res, temp)
		for i := u; i < len(nums); i++ {
			comb = append(comb, nums[i])
			dfs(i + 1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0)
	return res
}
