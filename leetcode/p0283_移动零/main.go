// https://leetcode-cn.com/problems/move-zeroes/submissions/

package main

func main() {

}

func moveZeroes(nums []int) {
	pos := 0
	for _, n := range nums {
		if n != 0 {
			nums[pos] = n
			pos++
		}
	}
	for pos < len(nums) {
		nums[pos] = 0
		pos++
	}
}
