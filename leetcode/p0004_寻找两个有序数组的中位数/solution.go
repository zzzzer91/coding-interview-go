package p0004_寻找两个有序数组的中位数

// 开辟一个辅助数组，通俗易懂，但时间复杂度无法达到 O(log(m+n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	tempArray := make([]int, totalLen)
	pos := 0
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			tempArray[pos] = nums1[i]
			i++
		} else {
			tempArray[pos] = nums2[j]
			j++
		}
		pos++
	}
	for i < len(nums1) {
		tempArray[pos] = nums1[i]
		i++
		pos++
	}
	for j < len(nums2) {
		tempArray[pos] = nums2[j]
		j++
		pos++
	}
	if totalLen%2 == 0 {
		return float64(tempArray[totalLen/2-1]+tempArray[totalLen/2]) / 2
	}
	return float64(tempArray[totalLen/2])
}
