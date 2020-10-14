// https://leetcode-cn.com/problems/intersection-of-two-arrays/

package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	m := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		m[num] = 0
	}
	for _, num := range nums2 {
		if v, ok := m[num]; ok && v == 0 {
			res = append(res, num)
			m[num]++
		}
	}
	return res
}
