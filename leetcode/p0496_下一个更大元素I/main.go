// https://leetcode-cn.com/problems/next-greater-element-i/

package main

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	m := make(map[int]int, len(nums1))
	var stack []int
	for _, n := range nums2 {
		for len(stack) > 0 && n > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = n
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, n)
	}
	for i, n := range nums1 {
		if v, ok := m[n]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}
	return res
}
