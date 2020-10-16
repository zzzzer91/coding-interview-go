// https://leetcode-cn.com/problems/merge-sorted-array/

package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	pos := m + n - 1
	// 如果 j 先结束，那么 nums1 剩下的元素其实不需要动了
	for j >= 0 {
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[pos] = nums1[i]
			i--
		} else {
			nums1[pos] = nums2[j]
			j--
		}
		pos--
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	pos := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[pos] = nums2[j]
			pos--
			j--
		} else {
			nums1[pos] = nums1[i]
			pos--
			i--
		}
	}
	for j >= 0 {
		nums1[pos] = nums2[j]
		pos--
		j--
	}
}
