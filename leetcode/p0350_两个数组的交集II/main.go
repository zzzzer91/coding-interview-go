// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/

package main

import "sort"

func main() {

}

// 使用字典
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}
	for _, num := range nums2 {
		if m[num] > 0 {
			res = append(res, num)
			m[num]--
		}
	}
	return res
}

// 假设两个数组已排好序
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}
