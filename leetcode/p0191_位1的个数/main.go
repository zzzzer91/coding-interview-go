// https://leetcode-cn.com/problems/number-of-1-bits/

package main

func main() {

}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num & 1)
		num >>= 1
	}
	return count
}

// 性能更好
func hammingWeight2(num uint32) int {
	count := 0
	for num != 0 {
		// 这个技巧能清0右边最后一个bit
		num &= num - 1
		count++
	}
	return count
}
