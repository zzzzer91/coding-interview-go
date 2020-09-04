package p0026_删除排序数组中的重复项

func removeDuplicates(nums []int) int {
	pos := 0
	for i := 1; i < len(nums); i++ {
		if nums[pos] != nums[i] {
			pos++
			nums[pos] = nums[i]
		}
	}
	return pos + 1
}
