// https://leetcode-cn.com/problems/contains-duplicate/

package main

func main() {

}

func containsDuplicate(nums []int) bool {
	intSet := make(map[int]bool)
	for _, v := range nums {
		if intSet[v] {
			return true
		}
		intSet[v] = true
	}
	return false
}
