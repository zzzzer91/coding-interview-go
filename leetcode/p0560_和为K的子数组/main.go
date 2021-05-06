// https://leetcode-cn.com/problems/subarray-sum-equals-k/

package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 2, 1, 2, 1}, 3))
}

// 枚举
func subarraySum(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// 前缀和 + 哈希表优化
func subarraySum2(nums []int, k int) int {
	res := 0
	m := make(map[int]int) // 和为键，次数为值
	m[0] = 1               // 考虑子数组和为 k 的情况
	prefixSum := 0
	for _, n := range nums {
		prefixSum += n
		res += m[prefixSum-k]
		m[prefixSum] += 1
	}
	return res
}
