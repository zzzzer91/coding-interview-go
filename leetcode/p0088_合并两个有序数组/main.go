package main

func main() {

}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	i, j := m-1, n-1
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
	pos := m + n - 1
	i, j := m-1, n-1
	for j >= 0 {
		if i >= 0 {
			if nums1[i] >= nums2[j] {
				nums1[pos] = nums1[i]
				i--
			} else {
				nums1[pos] = nums2[j]
				j--
			}
			pos--
		} else {
			// 用 copy 也许能提高效率？
			copy(nums1, nums2[:j+1])
			break
		}
	}
}
