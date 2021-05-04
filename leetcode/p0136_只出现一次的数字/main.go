// https://leetcode-cn.com/problems/single-number/

package main

func main() {

}

func singleNumber(nums []int) int {
	ret := 0
	for _, n := range nums {
		ret ^= n
	}
	return ret
}
