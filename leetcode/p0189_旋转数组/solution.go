package main

func main() {

}

func rotate(nums []int, k int) {
	numsLen := len(nums)
	if numsLen == 0 || k == 0 {
		return
	}
	k %= numsLen
	reverse(nums[:numsLen-k])
	reverse(nums[numsLen-k:])
	reverse(nums)
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
