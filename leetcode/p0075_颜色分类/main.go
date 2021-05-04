// https://leetcode-cn.com/problems/sort-colors/

package main

func main() {

}

// 三指针
// 多模拟几遍去理解
// [1,2,0]
func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	cur := 0
	for cur <= r { // 注意有等于条件
		if nums[cur] == 0 {
			nums[cur], nums[l] = nums[l], nums[cur]
			l++
			cur++
		} else if nums[cur] == 2 {
			nums[cur], nums[r] = nums[r], nums[cur]
			r-- // 没有 cur++
		} else {
			cur++
		}
	}
}
