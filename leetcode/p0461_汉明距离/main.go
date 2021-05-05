// https://leetcode-cn.com/problems/hamming-distance/

package main

func main() {

}

func hammingDistance(x int, y int) int {
	return bitCount(x ^ y)
}

func bitCount(x int) int {
	res := 0
	for x > 0 {
		x &= x - 1
		res++
	}
	return res
}
