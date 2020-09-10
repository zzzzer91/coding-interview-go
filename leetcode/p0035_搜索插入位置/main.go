package main

func main() {

}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		val := nums[mid]
		if target < val {
			right = mid - 1
		} else if target > val {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
