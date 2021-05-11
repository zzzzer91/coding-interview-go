// https://leetcode-cn.com/problems/longest-consecutive-sequence/

package main

import "sort"

func main() {

}

// 先排序，然后统计
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	sort.Ints(nums)
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			count++
		} else if nums[i] == nums[i-1] {
			continue
		} else if nums[i]-1 != nums[i-1] {
			res = max(res, count)
			count = 1
		}
	}
	res = max(res, count)
	return res
}

func longestConsecutive2(nums []int) int {
	res := 0
	m := make(map[int]bool, len(nums))
	for _, n := range nums {
		m[n] = true
	}
	for _, n := range nums {
		// 如果前面还有连续的，则跳过当前的
		if !m[n-1] {
			length := 1
			// 不停的往后找连续的
			for i := n + 1; m[i]; i++ {
				length++
			}
			// 与之前长度比对
			res = max(res, length)
		}
	}
	return res
}

func longestConsecutive3(nums []int) int {
	res := 0
	m := make(map[int]int)
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			left := m[n-1]
			right := m[n+1]
			curLength := 1 + left + right
			res = max(res, curLength)
			m[n] = curLength
			m[n-left] = curLength
			m[n+right] = curLength
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
