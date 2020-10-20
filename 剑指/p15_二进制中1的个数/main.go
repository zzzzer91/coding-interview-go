// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/

package main

func main() {

}

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		num &= num - 1
		res++
	}
	return res
}
