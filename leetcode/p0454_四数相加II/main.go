// https://leetcode-cn.com/problems/4sum-ii/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	res := 0
	m := make(map[int]int, len(nums1)*len(nums2)) // 提前分配好空间会快很多
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			m[-(n1+n2)]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			res += m[n3+n4]
		}
	}
	return res
}
