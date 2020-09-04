package p0167_两数之和II

func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1
	for lo < hi {
		sum := numbers[lo] + numbers[hi]
		if sum > target {
			hi--
		} else if sum < target {
			lo++
		} else {
			// 数组首元素索引为 1，所以要加一
			return []int{lo + 1, hi + 1}
		}
	}
	return nil
}
