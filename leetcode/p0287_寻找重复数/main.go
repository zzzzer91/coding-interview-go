// https://leetcode-cn.com/problems/find-the-duplicate-number/

package main

func main() {

}

// 快慢指针思想，参考判断链表有环
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	slow, fast = nums[slow], nums[nums[fast]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}
