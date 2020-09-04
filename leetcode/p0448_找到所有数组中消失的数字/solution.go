package p0448_找到所有数组中消失的数字

// 将出现过的元素作为下标，将数组中对应元素标记为负的，
// 若数组某元素为正的，说明其下标+1没在数组出现过，就是我们要找的值
func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		idx := abs(num) - 1
		nums[idx] = -abs(nums[idx])
	}
	var ret []int
	for i, num := range nums {
		if num >= 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
