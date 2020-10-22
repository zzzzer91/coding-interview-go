// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

package main

func main() {

}

func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]&1 != 0 {
			l++
		}
		for l < r && nums[r]&1 == 0 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
