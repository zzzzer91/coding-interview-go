package p0027_移除元素

func removeElement(nums []int, val int) int {
	pos := 0
	for _, num := range nums {
		if num != val {
			nums[pos] = num
			pos++
		}
	}
	return pos
}
