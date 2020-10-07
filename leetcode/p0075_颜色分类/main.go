// https://leetcode-cn.com/problems/sort-colors/

package main

func main() {

}

// 三指针
func sortColors(nums []int) {
	p, q := 0, len(nums)-1
	cur := 0
	for cur <= q {
		if nums[cur] == 0 {
			nums[cur], nums[p] = nums[p], nums[cur]
			p++
			cur++
		} else if nums[cur] == 2 {
			nums[cur], nums[q] = nums[q], nums[cur]
			q--
		} else {
			cur++
		}
	}
}
