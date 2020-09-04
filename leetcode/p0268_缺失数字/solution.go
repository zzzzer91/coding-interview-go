package p0268_缺失数字

func missingNumber(nums []int) int {
	n := len(nums) * (len(nums) + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return n - sum
}

// 方法2，用异或，感觉更好
func missingNumber2(nums []int) int {
	sum := 0
	for i, num := range nums {
		sum = sum ^ i ^ num
	}
	return sum ^ len(nums)
}
