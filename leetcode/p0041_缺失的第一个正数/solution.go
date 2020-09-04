package p0041_缺失的第一个正数

func firstMissingPositive(nums []int) int {
	for _, num := range nums {
		for num > 0 && num <= len(nums) && nums[num-1] != num {
			nums[num-1], num = num, nums[num-1]
		}
	}
	for i, num := range nums {
		if i+1 != num {
			return i + 1
		}
	}
	return len(nums) + 1
}
