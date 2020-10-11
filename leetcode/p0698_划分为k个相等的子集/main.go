// https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/

package main

import "sort"

func main() {

}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	sum /= k
	sort.Ints(nums)
	if nums[len(nums)-1] > sum {
		return false
	}
	buckets := make([]int, k)
	for i := range buckets {
		buckets[i] = sum
	}
	var dfs func(u int) bool
	dfs = func(u int) bool {
		if u == -1 { // 说明 nums 中元素都用完了，即所有元素都正确放在了桶中
			return true
		}
		for i := 0; i < k; i++ {
			// buckets[i]-nums[u] >= nums[0]，剪枝，桶的剩余容量必须大于等于元素中的最小值
			if buckets[i] == nums[u] || buckets[i]-nums[u] >= nums[0] {
				buckets[i] -= nums[u]
				if dfs(u - 1) {
					return true
				}
				buckets[i] += nums[u]
			}
		}
		return false
	}
	return dfs(len(nums) - 1)
}
