// https://leetcode-cn.com/problems/two-sum/
// 相似题目：
// https://leetcode-cn.com/problems/4sum-ii/
// https://leetcode-cn.com/submissions/detail/174215144/

package main

func main() {

}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, n := range nums {
		if v, ok := dict[n]; ok {
			return []int{v, i}
		}
		dict[target-n] = i
	}
	return []int{-1, -1}
}
