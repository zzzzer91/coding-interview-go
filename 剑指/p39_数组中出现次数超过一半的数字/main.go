// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/

package main

func main() {

}

// 摩尔投票法：核心理念为票数正负抵消。
// 此方法时间和空间复杂度分别为 O(N) 和 O(1)，为本题的最佳解法。
func majorityElement(nums []int) int {
	res := nums[0]
	count := 1
	for _, n := range nums[1:] {
		if n == res {
			count++
		} else {
			count--
			if count == 0 {
				res = n
				count = 1
			}
		}
	}
	return res
}
