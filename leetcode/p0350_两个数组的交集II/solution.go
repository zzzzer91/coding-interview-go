package p0350_两个数组的交集II

import "sort"

// 使用字典
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}
	var ret []int
	for _, num := range nums2 {
		if m[num] > 0 {
			ret = append(ret, num)
			m[num]--
		}
	}
	return ret
}

// 假设两个数组已排好序
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var ret []int
	i, j := 0, 0
	for i != len(nums1) && j != len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			ret = append(ret, nums1[i])
			i++
			j++
		}
	}
	return ret
}
